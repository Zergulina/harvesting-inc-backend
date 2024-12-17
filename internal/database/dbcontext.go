package database

import (
	"backend/internal/config"
	"backend/internal/database/repository"
	"backend/internal/helpers"
	"backend/internal/models"
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	dbConn, err := sql.Open("postgres", config.DbConnectionString)
	if err != nil {
		return err
	}

	DB = dbConn

	err = DB.Ping()
	if err != nil {
		panic("Error: Unable to ping database")
	}

	err = initDb(dbConn)
	if err != nil {
		panic(err)
	}

	return nil
}

func initDb(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS people (
			id SERIAL NOT NULL PRIMARY KEY,
			lastname TEXT NOT NULL,
			firstname TEXT NOT NULL,
			middlename TEXT,
			birthdate DATE NOT NULL,
			login TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT UNIQUE
		);

		CREATE TABLE IF NOT EXISTS employees (
			people_id INTEGER NOT NULL REFERENCES people(id) ON DELETE CASCADE,
			post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
			employment_date DATE NOT NULL,
			fire_date DATE,
			salary INTEGER NOT NULL,
			PRIMARY KEY(people_id, post_id)
		);
		
		CREATE TABLE IF NOT EXISTS customers (
			id SERIAL NOT NULL PRIMARY KEY,
			ogrn CHAR(13) NOT NULL,
			name TEXT NOT NULL,
			logo BYTEA,
			logo_extension TEXT
		);

		CREATE TABLE IF NOT EXISTS crop_types (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS crops (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			crop_type_id INTEGER NOT NULL REFERENCES crop_types(id) ON DELETE CASCADE,
			description TEXT
		);

		CREATE TABLE IF NOT EXISTS fields (
			id SERIAL NOT NULL PRIMARY KEY,
			coords TEXT NOT NULL,
			customer_id INTEGER NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
			crop_id INTEGER NOT NULL REFERENCES crops(id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS statuses (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			is_available BOOLEAN NOT NULL
		);

		CREATE TABLE IF NOT EXISTS machine_types (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS machine_models (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			machine_type_id INTEGER NOT NULL REFERENCES machine_types(id) ON DELETE CASCADE
		);
		
		CREATE TABLE IF NOT EXISTS machines (
			inv_number INTEGER NOT NULL,
			machine_model_id INTEGER NOT NULL REFERENCES machine_models(id) ON DELETE CASCADE,
			status_id INTEGER NOT NULL REFERENCES statuses(id) ON DELETE CASCADE,
			buy_date DATE NOT NULL,
			draw_down_date DATE,
			PRIMARY KEY(inv_number, machine_model_id)
		);

		CREATE TABLE IF NOT EXISTS equipment_types (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS equipment_models (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			equipment_type_id INTEGER NOT NULL REFERENCES equipment_types(id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS equipments (
			inv_number INTEGER NOT NULL,
			equipment_model_id INTEGER NOT NULL REFERENCES equipment_models(id) ON DELETE CASCADE,
			status_id INTEGER NOT NULL REFERENCES statuses(id) ON DELETE CASCADE,
			buy_date DATE NOT NULL,
			draw_down_date DATE,
			PRIMARY KEY(inv_number, equipment_model_id)
		);

		CREATE TABLE IF NOT EXISTS machine_equipment_types (
			machine_type_id INTEGER NOT NULL REFERENCES machine_types(id) ON DELETE CASCADE,
			equipment_type_id INTEGER NOT NULL REFERENCES equipment_types(id) ON DELETE CASCADE,
			PRIMARY KEY(machine_type_id, equipment_type_id)
		);

		CREATE TABLE IF NOT EXISTS works (
			id SERIAL NOT NULL PRIMARY KEY,
			start_date DATE NOT NULL,
			end_date DATE,
			field_id INTEGER NOT NULL REFERENCES fields(id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS work_trips (
			id SERIAL NOT NULL PRIMARY KEY,
			start_date TIMESTAMP NOT NULL,
			end_date TIMESTAMP,
			people_id INTEGER NOT NULL REFERENCES people(id) ON DELETE CASCADE,
			crop_amount INTEGER NOT NULL,
			work_id INTEGER NOT NULL REFERENCES works(id) ON DELETE CASCADE,
			machine_inv_number INTEGER NOT NULL,
			machine_model_id INTEGER NOT NULL,
			FOREIGN KEY (machine_inv_number, machine_model_id) REFERENCES machines(inv_number, machine_model_id) ON DELETE CASCADE,
			equipment_inv_number INTEGER,
			equipment_model_id INTEGER,
			FOREIGN KEY (equipment_inv_number, equipment_model_id) REFERENCES equipments(inv_number, equipment_model_id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS vacations (
			people_id INTEGER NOT NULL REFERENCES people(id) ON DELETE CASCADE,
			start_date DATE NOT NULL,
			end_date DATE NOT NULL,
			PRIMARY KEY(people_id, start_date)
		);
		`)

	if err != nil {
		return err
	}

	var adminPostId uint64

	isExists, err := repository.ExistsPostByName(db, config.AdminRole)
	if err != nil {
		return err
	}
	if !isExists {
		post, err := repository.CreatePost(db, &models.Post{Name: config.AdminRole})
		if err != nil {
			return err
		}

		adminPostId = post.Id
	} else {
		adminPost, err := repository.GetPostByName(db, config.AdminRole)
		if err != nil {
			return err
		}

		adminPostId = adminPost.Id
	}

	isExists, err = repository.ExistsPostByName(db, config.HrRole)
	if err != nil {
		return err
	}
	if !isExists {
		_, err = repository.CreatePost(db, &models.Post{Name: config.HrRole})
		if err != nil {
			return err
		}
	}

	isExists, err = repository.ExistsPostByName(db, config.DriverRole)
	if err != nil {
		return err
	}
	if !isExists {
		_, err = repository.CreatePost(db, &models.Post{Name: config.DriverRole})
		if err != nil {
			return err
		}
	}

	var adminId uint64

	isExists, err = repository.ExistsPeopleByLogin(db, config.AdminLogin)
	if err != nil {
		return err
	}
	if !isExists {
		admin := new(models.People)
		admin = &models.People{
			LastName:     config.AdminLastname,
			FirstName:    config.AdminFirstname,
			MiddleName:   config.AdminMiddlename,
			BirthDate:    config.AdminBirthdate,
			Login:        config.AdminLogin,
			PasswordHash: helpers.EncodeSha256(config.AdminPassword, config.DbSecretKey),
		}
		admin, err = repository.CreatePeople(db, admin)
		if err != nil {
			return err
		}
		adminId = admin.Id
	} else {
		admin, err := repository.GetPeopleByLogin(db, config.AdminLogin)
		if err != nil {
			return err
		}
		adminId = admin.Id
	}

	isExists, err = repository.ExistsEmployee(db, adminId, adminPostId)
	if err != nil {
		return err
	}
	if !isExists {
		employee := new(models.Employee)
		employee = &models.Employee{
			PeopleId:       adminId,
			PostId:         adminPostId,
			EmploymentDate: config.AdminEmploymentDate,
			Salary:         config.AdminSalary,
		}
		_, err = repository.CreateEmployee(db, employee)
		if err != nil {
			return err
		}
	}

	return nil
}

func ResetDb() {
	_, err := DB.Exec(`
			DROP SCHEMA public CASCADE;
			CREATE SCHEMA public;

			CREATE TABLE IF NOT EXISTS people (
				id SERIAL NOT NULL PRIMARY KEY,
				lastname TEXT NOT NULL,
				firstname TEXT NOT NULL,
				middlename TEXT,
				birthdate DATE NOT NULL,
				login TEXT NOT NULL UNIQUE,
				password_hash TEXT NOT NULL
			);
			
			CREATE TABLE IF NOT EXISTS posts (
				id SERIAL NOT NULL PRIMARY KEY,
				name TEXT UNIQUE
			);

			CREATE TABLE IF NOT EXISTS employees (
				people_id INTEGER NOT NULL REFERENCES people(id) ON DELETE CASCADE,
				post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
				employment_date DATE NOT NULL,
				fire_date DATE,
				salary INTEGER NOT NULL,
				PRIMARY KEY(people_id, post_id)
			);
			
			CREATE TABLE IF NOT EXISTS customers (
				id SERIAL NOT NULL PRIMARY KEY,
				ogrn CHAR(13) NOT NULL,
				name TEXT NOT NULL,
				logo BYTEA,
				logo_extension TEXT
			);

			CREATE TABLE IF NOT EXISTS crop_types (
				id SERIAL NOT NULL PRIMARY KEY,
				name TEXT NOT NULL
			);

			CREATE TABLE IF NOT EXISTS crops (
				id SERIAL NOT NULL PRIMARY KEY,
				name TEXT NOT NULL,
				crop_type_id INTEGER NOT NULL REFERENCES crop_types(id) ON DELETE CASCADE,
				description TEXT
			);

			CREATE TABLE IF NOT EXISTS fields (
				id SERIAL NOT NULL PRIMARY KEY,
				coords TEXT NOT NULL,
				customer_id INTEGER NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
				crop_id INTEGER NOT NULL REFERENCES crops(id) ON DELETE CASCADE
			);

			CREATE TABLE IF NOT EXISTS statuses (
				id SERIAL NOT NULL PRIMARY KEY,
				name TEXT NOT NULL,
				is_available BOOLEAN NOT NULL
			);

			CREATE TABLE IF NOT EXISTS machine_types (
				id SERIAL NOT NULL PRIMARY KEY,
				name TEXT NOT NULL
			);

			CREATE TABLE IF NOT EXISTS machine_models (
				id SERIAL NOT NULL PRIMARY KEY,
				name TEXT NOT NULL,
				machine_type_id INTEGER NOT NULL REFERENCES machine_types(id) ON DELETE CASCADE
			);
			
			CREATE TABLE IF NOT EXISTS machines (
				inv_number INTEGER NOT NULL,
				machine_model_id INTEGER NOT NULL REFERENCES machine_models(id) ON DELETE CASCADE,
				status_id INTEGER NOT NULL REFERENCES statuses(id) ON DELETE CASCADE,
				buy_date DATE NOT NULL,
				draw_down_date DATE,
				PRIMARY KEY(inv_number, machine_model_id)
			);

			CREATE TABLE IF NOT EXISTS equipment_types (
				id SERIAL NOT NULL PRIMARY KEY,
				name TEXT NOT NULL
			);

			CREATE TABLE IF NOT EXISTS equipment_models (
				id SERIAL NOT NULL PRIMARY KEY,
				name TEXT NOT NULL,
				equipment_type_id INTEGER NOT NULL REFERENCES equipment_types(id) ON DELETE CASCADE
			);

			CREATE TABLE IF NOT EXISTS equipments (
				inv_number INTEGER NOT NULL,
				equipment_model_id INTEGER NOT NULL REFERENCES equipment_models(id) ON DELETE CASCADE,
				status_id INTEGER NOT NULL REFERENCES statuses(id) ON DELETE CASCADE,
				buy_date DATE NOT NULL,
				draw_down_date DATE,
				PRIMARY KEY(inv_number, equipment_model_id)
			);

			CREATE TABLE IF NOT EXISTS machine_equipment_types (
				machine_type_id INTEGER NOT NULL REFERENCES machine_types(id) ON DELETE CASCADE,
				equipment_type_id INTEGER NOT NULL REFERENCES equipment_types(id) ON DELETE CASCADE,
				PRIMARY KEY(machine_type_id, equipment_type_id)
			);

			CREATE TABLE IF NOT EXISTS works (
				id SERIAL NOT NULL PRIMARY KEY,
				start_date DATE NOT NULL,
				end_date DATE,
				field_id INTEGER NOT NULL REFERENCES fields(id) ON DELETE CASCADE
			);

			CREATE TABLE IF NOT EXISTS work_trips (
				id SERIAL NOT NULL PRIMARY KEY,
				start_date TIMESTAMP NOT NULL,
				end_date TIMESTAMP,
				people_id INTEGER NOT NULL REFERENCES people(id) ON DELETE CASCADE,
				crop_amount INTEGER NOT NULL,
				work_id INTEGER NOT NULL REFERENCES works(id) ON DELETE CASCADE,
				machine_inv_number INTEGER NOT NULL,
				machine_model_id INTEGER NOT NULL,
				FOREIGN KEY (machine_inv_number, machine_model_id) REFERENCES machines(inv_number, machine_model_id) ON DELETE CASCADE,
				equipment_inv_number INTEGER,
				equipment_model_id INTEGER,
				FOREIGN KEY (equipment_inv_number, equipment_model_id) REFERENCES equipments(inv_number, equipment_model_id) ON DELETE CASCADE
			);

			CREATE TABLE IF NOT EXISTS vacations (
				people_id INTEGER NOT NULL REFERENCES people(id) ON DELETE CASCADE,
				start_date DATE NOT NULL,
				end_date DATE NOT NULL,
				PRIMARY KEY(people_id, start_date)
			);
	        `)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO people (lastname, firstname, middlename, birthdate, login, password_hash)
				VALUES 
				($1, $2, $3, $4, $5, $6),
				('Иванов', 'Иван', 'Иванович', '1990-01-01', 'ivanov.ii', $7),
				('Петров', 'Пётр', 'Петрович', '1985-06-15', 'petrov.pp', $8),
				('Сидоров', 'Сидор', 'Сидорович', '1992-02-20', 'sidorov.ss', $9),
				('Табуретка', 'Нампай', 'Сервлетович', '2001-02-01', 'kocherga.va', $10);
			`,
		config.AdminLastname,
		config.AdminFirstname,
		config.AdminMiddlename,
		config.AdminBirthdate,
		config.AdminLogin,
		helpers.EncodeSha256(config.AdminPassword, config.DbSecretKey),
		helpers.EncodeSha256("12345678", config.DbSecretKey),
		helpers.EncodeSha256("12345678", config.DbSecretKey),
		helpers.EncodeSha256("12345678", config.DbSecretKey),
		helpers.EncodeSha256("12345678", config.DbSecretKey))
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO posts (name)
				VALUES 
				($1),
				($2),
				($3);
			`, config.AdminRole, config.HrRole, config.DriverRole)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO employees (people_id, post_id, employment_date, fire_date, salary)
				VALUES 
				(1, 1, $1, NULL, $2),
				(1, 2, '2010-01-02', NULL, 4000000),
				(2, 2, '2011-01-01', NULL, 5500000),
				(3, 3, '2015-06-15', NULL, 4500000),
				(4, 2, '2020-01-11', NULL, 6000000),
				(5, 2, '2021-12-23', NULL, 5000000),
				(5, 3, '2022-05-05', NULL, 3000000);
			`, config.AdminEmploymentDate, config.AdminSalary)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO customers (ogrn, name)
				VALUES 
				('1234567890123', 'ОПХ "Садовод"'),
				('9876543210987', 'ЗАО "Русская ферма"'),
				('5678901234567', 'ПАО "Фермерское хозяйство"'),
				('1147746518017', 'Агрохолдинг "Степь"'),
				('1147746518013', 'Агрохолдинг "Качаван"');
			`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO crop_types (name)
				VALUES 
				('Пшеница'),
				('Рожь'),
				('Ячмень'),
				('Подсолнечник'),
				('Кукуруза');
			`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO crops (name, crop_type_id, description)
				VALUES 
				('Антоновка', 1, 'Высота растений не превышает 95 см, колоски белого цвета, без признаков опушения. Приспосабливается к разным погодным условиям, устойчивый к засухам и многим распространённым заболеваниям.'),
				('Рино', 1, 'Высота растений превышает 1 м, не подвержен осыпанию, вынослив к неблагоприятным погодным условиям, устойчив к наиболее распространённым заболеваниям. Созревание происходит на протяжении 280 дней.'),
				('Чикаго', 1, 'Зёрна крупные, их масса может достичь 50 г (1000 шт.). Не теряет качеств на протяжении 8 лет высевания, не подвержен осыпанию, устойчив к большому количеству заболеваний. Для полного созревания требуется не менее 300 дней.'),
				('Василина', 1, 'Масса 1000 зёрен — 40 г, высота растений равна 90 см, зимостойкость оценивается в 9 баллов, урожайность составляет 90 ц/га. Созревает за период до 300 суток.'),
				('Харус', 1, 'Вес 1000 зёрен составляет примерно 45 г, высота растений не превышает 90 см, зимостойкость достаточно высокая и достигает 8 баллов, урожайность составляет около 80 ц/га.'),
				('Альфа', 2, 'Высота злака — от 115 до 120 см, зерно полукруглой формы, опушённое основание. Морозоустойчивая, позднеспелая — до 350 дней.'),
				('Эстафета Татарстана', 2, 'Высота растения — 176 см, форма зерна полуудлинённая, структура полуоткрытая. Созревание — до 331 дня.'),
				('Татьяна', 2, 'Длина стебля — 142 см, зёрна мелкие. Созревает за 349 дней, морозостойкая.'),
				('Радонь', 2, 'Высота стебля — 130 см, зерно крупного размера. Среднепозднее созревание — до 335 дней.'),
				('Верасень', 2, 'Растёт до 140 см в высоту. Созревает за 305–311 дней. Достоинства: устойчивость к полеганию, холодам, повышенной влажности или засухе. Зёрна крупные, удлинённой формы.'),
				('Виконт', 3, 'Гибридный сорт, растение прямостоячее, созревает за 90 дней после сева. Вес 1000 зёрен — примерно 50–80 грамм. Чаще всего используется для пивоварения.'),
				('Приазовский', 3, 'Самый распространённый сорт на территории РФ. Имеет высокую жизнестойкость и может дать хороший урожай даже на необогащённом грунте. Срок созревания — 90 дней, имеет высокую устойчивость к полеганию, холодостойкий.'),
				('Гелиос', 3, 'Сорт неприхотлив к качеству грунта и имеет высокую всхожесть. При повышенной влажности может дать хороший урожай зерна. Срок созревания — 90 дней. '),
				('Мамлюк', 3, 'У сорта высокая всхожесть, он раннеспелый и очень продуктивный. Ему не страшна непродолжительная засуха, а ещё он устойчив перед большинством форм грибка. Чаще всего данный сорт выращивают на фураж и для переработки на крупу.'),
				('Вакула', 3, 'Сорт даёт хороший урожай, так как он максимально адаптирован к климатическим изменениям. С 1 га засеянного поля можно получить примерно 85 центнеров ячменя.'),
				('Сузука', 4, 'Среднеранний, предназначен для засушливых условий возделывания.'),
				('НК Неома', 4, 'Высокоинтенсивный, среднеспелый гибрид подсолнечника для производственной системы Clearfield®. Масличность 50–52%.'),
				('П 63 ЛЕ 10', 4, 'Раннеспелый, гибрид нового поколения, устойчив к заразихе. Среднее содержание жира в семенах 48,5%.'),
				('Кречет', 4, 'Среднеспелый, обладает высокой стрессоустойчивостью. Максимальная урожайность 38,7 ц/га получена на Отрадненском ГСУ Краснодарского края в 2018 году.'),
				('Донской крупноплодный', 4, 'Среднеспелый сорт, формирующий крупную семянку с высоким качеством ядра.'),
				('Жемчуг', 5, 'Урожай можно получить через 83–91 день. Вырастает до 1,7 м, початки формируются длиной около 21 см. Зёрна хорошо хранятся после уборки.'),
				('Свитстар F1', 5, 'Сладкий гибрид со сроком вегетации около 73 дней. Вырастает до 2,2 м, початки довольно крупные — около 24 см. Зёрна имеют жёлтую окраску.'),
				('Спирит F1', 5, 'Гибрид созревает за 63–79 дней, вырастает до 2 м. Можно получить початки длиной около 23 см, которые имеют сочные зёрна золотистого цвета с высоким содержанием сахара.'),
				('Фаворит', 5, 'Потребуется 58–66 дней, чтобы кукуруза созрела. Вырастает до 1,8 см. Початки имеют длину около 23 см и зёрна ярко-жёлтого цвета.'),
				('Деликатесная', 5, 'Сахарная кукуруза со сроком вегетации до 70 дней. Высота растения небольшая — до 1,4 м, но початок вырастает до 22 см. Зёрна насыщенного жёлтого цвета и хорошо поддаются любой переработке.');
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO fields (coords, customer_id, crop_id)
				VALUES 
				('52.1234, 45.5678', 1, 1),
				('54.9012, 43.1111', 1, 12),
				('53.4567, 42.7890', 2, 3),
				('52.1234, 45.5678', 2, 4),
				('54.9054, 43.1111', 3, 5),
				('53.2434, 42.7890', 3, 6),
				('52.6543, 45.5678', 4, 7),
				('55.1234, 43.1111', 4, 8),
				('50.4567, 42.7890', 5, 9),
				('49.1234, 45.5678', 5, 10),
				('51.9012, 43.1111', 5, 11),
				('51.9967, 42.7890', 5, 12);
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO statuses (name, is_available)
				VALUES 
				('На выезде', FALSE),
				('В ремонте', FALSE),
				('Списан', FALSE),
				('На выезде', FALSE),
				('В ангаре', TRUE);
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO machine_types (name)
				VALUES 
				('Гусеничный трактор'),
				('Колесный трактор'),
				('Роторный комбайн'),
				('Клавишный комбайн'),
				('Зерновоз');
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO machine_models (name, machine_type_id)
				VALUES 
				('Steiger QuadTrac 385', 1),
				('Versatile Deltatrack 520DT', 1),
				('АЛТТРАК А600', 1),
				('МТЗ 82.1', 2),
				('МТЗ 1221', 2),
				('YTO 2204', 2),
				('TORUM 785', 3),
				('NEW HOLLAND CX 8', 4),
				('MАЗ 6312C9-8575-012', 5);
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO machines (inv_number, machine_model_id, status_id, buy_date, draw_down_date)
				VALUES 
				(1, 1, 5, '2010-01-01', NULL),
				(2, 1, 5, '2010-01-01', NULL),
				(1, 4, 5, '2010-01-01', NULL),
				(1, 7, 5, '2010-01-01', NULL),
				(1, 8, 5, '2015-06-15', NULL),
				(1, 9, 5, '2020-02-20', NULL);
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO equipment_types (name)
				VALUES 
				('Самосвальный полуприцеп'),
				('Комбайн зерноуборочный прицепной'),
				('Прицеп'),
				('Жатка'),
				('Бункер-перегрузчик');
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO equipment_models (name, equipment_type_id)
				VALUES 
				('ППТС-5', 1),
				('ПСТ-6', 1),
				('ПН-100', 2),
				('Палессе 2U250A', 2),
				('Палессе 2U280A', 2),
				('Unicorn ЖСК-7', 4),
				('Liliani БП-33/42', 5);
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO equipments (inv_number, equipment_model_id, status_id, buy_date, draw_down_date)
				VALUES 
				(1, 1, 5, '2010-01-01', NULL),
				(2, 1, 5, '2010-06-15', NULL),
				(1, 2, 5, '2010-01-01', NULL),
				(1, 3, 5, '2010-06-15', NULL),
				(1, 4, 5, '2010-01-01', NULL),
				(1, 5, 5, '2010-06-15', NULL),
				(1, 6, 5, '2010-01-01', NULL),
				(1, 7, 5, '2010-06-15', NULL),
				(2, 4, 5, '2010-01-01', NULL),
				(2, 5, 5, '2010-06-15', NULL);
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO machine_equipment_types (machine_type_id, equipment_type_id)
				VALUES 
				(1, 1),
				(2, 1),
				(1, 2),
				(2, 2),
				(1, 5),
				(2, 5),
				(3, 4),
				(4, 4),
				(5, 3);
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO works (start_date, end_date, field_id)
				VALUES 
				('2024-06-01', NULL, 1),
				('2024-06-02', NULL, 2),
				('2024-07-20', NULL, 3),			
				('2024-07-21', NULL, 4),
				('2024-07-21', NULL, 5);
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO work_trips (start_date, end_date, people_id, crop_amount, work_id, machine_inv_number, machine_model_id, equipment_inv_number, equipment_model_id)
				VALUES 
				('2024-06-02', '2024-06-02', 4, 4, 1, 1, 1, NULL, NULL),
				('2024-06-02', '2024-06-02', 5, 4, 1, 2, 1, NULL, NULL),
				('2024-06-03', '2024-06-03', 4, 3, 1, 1, 1, NULL, NULL),
				('2024-06-03', '2024-06-03', 5, 2, 1, 2, 1, NULL, NULL),
				('2024-06-05', '2024-06-05', 4, 5, 1, 1, 1, NULL, NULL),
				('2024-06-05', '2024-06-05', 5, 1, 1, 2, 1, NULL, NULL),
				('2024-06-02', '2024-06-03', 4, 4, 2, 1, 1, NULL, NULL),
				('2024-06-02', '2024-06-03', 5, 1, 2, 2, 1, NULL, NULL),
				('2024-06-03', '2024-06-04', 4, 3, 2, 1, 1, NULL, NULL),
				('2024-06-03', '2024-06-04', 5, 2, 2, 2, 1, NULL, NULL),
				('2024-06-04', '2024-06-04', 5, 2, 2, 1, 1, NULL, NULL),
				('2024-06-04', '2024-06-04', 4, 1, 2, 2, 1, NULL, NULL)
				`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
			INSERT INTO vacations (people_id, start_date, end_date)
				VALUES 
				(1, '2020-07-15', '2020-08-01'),
				(2, '2015-12-20', '2016-01-05');
			`)
	if err != nil {
		panic(err)
	}
}
