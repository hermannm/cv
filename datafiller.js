window.onload = () => {
  const profilePicture = document.getElementById("profilepicture");
  const personFields = document.getElementById("person");
  const educationFields = document.getElementById("education");
  const experienceFields = document.getElementById("experience");

  profilePicture.src = data.person.profilepicture;

  Object.keys(data.person)
    .filter((field) => field !== "profilepicture")
    .map((field) => {
      const personField = document.createElement("li");
      personField.innerHTML += data.person[field].text;
      if (field === "name") {
        personField.className = "title";
      }
      personFields.appendChild(personField);
    });

  data.education.map((school) => {
    const educationField = document.createElement("li");
    educationField.className = "row spaced";

    const educationImg = document.createElement("img");
    educationImg.src = school.img;
    educationField.appendChild(educationImg);

    const educationText = document.createElement("div");
    educationField.appendChild(educationText);

    Object.keys(school)
      .filter((key) => key !== "img")
      .map((key) => {
        const child = document.createElement("div");
        child.innerHTML +=
          (key == "specialization"
            ? "Specialization: "
            : key == "location"
            ? ", "
            : "") + school[key];
        educationText.appendChild(child);
      });

    educationFields.appendChild(educationField);
  });
  data.experience.map((job) => {
    const experienceField = document.createElement("li");
    experienceField.className = "row spaced";

    const experienceImg = document.createElement("img");
    experienceImg.src = job.img;
    experienceField.appendChild(experienceImg);

    const experienceText = document.createElement("div");
    experienceField.appendChild(experienceText);

    Object.keys(job)
      .filter((field) => field !== "img")
      .map((field) => {
        const jobElement = document.createElement("div");
        if (field == "reference") {
          Object.keys(job[field]).map((key) => {
            const referenceElement = document.createElement("div");
            referenceElement.innerHTML += job[field][key];
            jobElement.appendChild(referenceElement);
          });
        } else {
          jobElement.innerHTML += job[field];
        }
        experienceText.appendChild(jobElement);
      });

    experienceFields.appendChild(experienceField);
  });
};
