package cvbuilder

import "html/template"

type CVTemplate struct {
	CV
	PersonalInfo PersonalInfo
}

type CV struct {
	EducationHeader              string           `yaml:"educationHeader"`
	Education                    []Education      `yaml:"education"`
	EducationSpecializationLabel string           `yaml:"educationSpecializationLabel"`
	WorkExperienceHeader         string           `yaml:"workExperienceHeader"`
	WorkExperience               []WorkExperience `yaml:"workExperience"`
	WorkExperienceReferenceLabel string           `yaml:"workExperienceReferenceLabel"`
}

type Education struct {
	StudyProgram   string `yaml:"studyProgram"`
	School         string `yaml:"school"`
	Time           string `yaml:"time"`
	Specialization string `yaml:"specialization"`
	ImagePath      string `yaml:"imagePath"`
}

type WorkExperience struct {
	Title        string                  `yaml:"title"`
	Organization template.HTML           `yaml:"organization"`
	Time         string                  `yaml:"time"`
	Description  template.HTML           `yaml:"description"`
	ImagePath    string                  `yaml:"imagePath"`
	Reference    WorkExperienceReference `yaml:"reference"` // Optional.
}

type WorkExperienceReference struct {
	Name        string `yaml:"name"`
	Title       string `yaml:"title"`
	PhoneNumber string `yaml:"phoneNumber"` // Optional.
	Email       string `yaml:"email"`       // Optional.
}

type JobApplicationTemplate struct {
	Application  template.HTML
	PersonalInfo PersonalInfo
}

type PersonalInfo struct {
	Name        string `yaml:"name"`
	Email       string `yaml:"email"`
	PhoneNumber string `yaml:"phoneNumber"`

	ProfilePicturePath string `yaml:"profilePicturePath"`
	SignaturePath      string `yaml:"signaturePath"`
	SignaturePrefix    string `yaml:"signaturePrefix"`

	Website struct {
		Text string `yaml:"text"`
		Link string `yaml:"link"`
	} `yaml:"website"` // Optional.
	GitHubLink   string `yaml:"githubLink"`   // Optional.
	LinkedInLink string `yaml:"linkedinLink"` // Optional.

	Age       string `yaml:"age"`       // If not set, Birthday and AgeSuffix must be set.
	Birthday  string `yaml:"birthday"`  // Required if Age is not set. Format: YYYY-MM-DD.
	AgeSuffix string `yaml:"ageSuffix"` // Required if Age is not set.
}
