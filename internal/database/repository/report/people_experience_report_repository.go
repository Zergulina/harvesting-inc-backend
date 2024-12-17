package report

import (
	"backend/internal/reports"
	"database/sql"
)

func GetPeopleExperienceReport(db *sql.DB) ([]reports.PeopleExperienceReport, error) {
	rows, err := db.Query(`WITH employee_experience AS 
							(SELECT people.id AS people_id, CONCAT_WS(' ', people.lastname, people.firstname, COALESCE(people.middlename, '')) AS fcs, posts.id AS post_id, posts.name AS post, EXTRACT (YEAR FROM AGE (COALESCE (fire_date, CURRENT_DATE), employees.employment_date)) AS experience 
							FROM people JOIN employees ON people.id = employees.people_id JOIN posts ON posts.id = employees.post_id)
						SELECT people_id, fcs, STRING_AGG(post, ', ') AS posts, SUM(experience) AS experience FROM employee_experience GROUP BY (people_id, fcs)`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	peopleExps := []reports.PeopleExperienceReport{}

	for rows.Next() {
		p := reports.PeopleExperienceReport{}
		err := rows.Scan(&p.PeopleId, &p.Fcs, &p.Posts, &p.Experience)
		if err != nil {
			continue
		}
		peopleExps = append(peopleExps, p)
	}

	return peopleExps, nil
}
