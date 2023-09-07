package cvbuilder

import (
	"fmt"
	"html/template"
	"time"

	"hermannm.dev/wrap"
)

type CVTemplate struct {
	CV
	PersonalInfo PersonalInfo
}

type CV struct {
	EducationHeader              string      `yaml:"educationHeader" validate:"required"`
	Education                    []Education `yaml:"education" validate:"required,dive"`
	EducationSpecializationLabel string      `yaml:"educationSpecializationLabel" validate:"required"`

	WorkExperienceHeader         string           `yaml:"workExperienceHeader" validate:"required"`
	WorkExperience               []WorkExperience `yaml:"workExperience" validate:"required,dive"`
	WorkExperienceReferenceLabel string           `yaml:"workExperienceReferenceLabel" validate:"required"`
}

type Education struct {
	StudyProgram   string `yaml:"studyProgram" validate:"required"`
	School         string `yaml:"school" validate:"required"`
	Time           string `yaml:"time" validate:"required"`
	Specialization string `yaml:"specialization" validate:"required"`
	ImagePath      string `yaml:"imagePath" validate:"required,filepath"`
}

type WorkExperience struct {
	Title        string                  `yaml:"title" validate:"required"`
	Organization template.HTML           `yaml:"organization" validate:"required"`
	Time         string                  `yaml:"time" validate:"required"`
	Description  template.HTML           `yaml:"description" validate:"required"`
	ImagePath    string                  `yaml:"imagePath" validate:"required,filepath"`
	Reference    WorkExperienceReference `yaml:"reference" validate:"excluded_without=Reference.Name"` // Optional.
}

type WorkExperienceReference struct {
	Name        string `yaml:"name"`
	Title       string `yaml:"title" validate:"required_with=Name"`
	PhoneNumber string `yaml:"phoneNumber"`
	Email       string `yaml:"email" validate:"omitempty,email"`
}

type JobApplicationTemplate struct {
	Application  template.HTML
	PersonalInfo PersonalInfo
}

type PersonalInfo struct {
	Name        string `yaml:"name" validate:"required"`
	Email       string `yaml:"email" validate:"required,email"`
	PhoneNumber string `yaml:"phoneNumber" validate:"required"`

	ProfilePicturePath string `yaml:"profilePicturePath" validate:"required,filepath"`
	SignatureImagePath string `yaml:"signatureImagePath" validate:"required,filepath"`
	SignaturePrefix    string `yaml:"signaturePrefix" validate:"required"`

	Website struct {
		Text string `yaml:"text" validate:"required_with=Link"`
		Link string `yaml:"link" validate:"url,required_with=Text"`
	} `yaml:"website"` // Optional.
	GitHubLink   string `yaml:"githubLink" validate:"url"`   // Optional.
	LinkedInLink string `yaml:"linkedinLink" validate:"url"` // Optional.

	Age       string `yaml:"age"` // Optional only if Birthday and AgeSuffix are set.
	Birthday  string `yaml:"birthday" validate:"required_without=Age,omitempty,datetime=2006-01-02"`
	AgeSuffix string `yaml:"ageSuffix" validate:"required_without=Age"`
}

func (info *PersonalInfo) setAge() error {
	if info.Age != "" {
		return nil
	}

	birthday, err := time.Parse(time.DateOnly, info.Birthday)
	if err != nil {
		return wrap.Error(err, "invalid format of birthday in personal info")
	}

	now := time.Now()
	age := now.Year() - birthday.Year()

	birthdayCelebratedThisYear := now.YearDay() >= birthday.YearDay()
	if !birthdayCelebratedThisYear {
		age--
	}

	info.Age = fmt.Sprintf("%d %s", age, info.AgeSuffix)
	return nil
}
