const getStyleClass = (key) => {
  switch (key) {
    case "time":
      return "bold";
    case "name":
    case "title":
      return "title";
    case "school":
    case "organization":
      return "gray";
    default:
      return "";
  }
};

window.onload = () => {
  const profilePicture = document.getElementById("profilepicture");
  const personFields = document.getElementById("person");
  const educationSectionTitle = document.getElementById("educationSection");
  const educationFields = document.getElementById("education");
  const experienceSectionTitle = document.getElementById("experienceSection");
  const experienceFields = document.getElementById("experience");

  educationSectionTitle.textContent = norwegian ? "Utdanning" : "Education";
  experienceSectionTitle.textContent = norwegian ? "Erfaring" : "Experience";

  profilePicture.src = data.person.profilepicture;

  Object.keys(data.person)
    .filter((field) => field !== "profilepicture")
    .map((field) => {
      const personField = document.createElement("li");
      if (data.person[field].link) {
        const linkField = document.createElement("a");
        linkField.textContent = data.person[field].text;
        linkField.href = data.person[field].link;
        personField.appendChild(linkField);
      } else {
        personField.textContent += data.person[field].text;
      }
      personField.className = `textField ${getStyleClass(field)}`;
      personFields.appendChild(personField);
    });

  data.education.map((school) => {
    const educationField = document.createElement("li");
    educationField.className = "row spaced dataField";

    const educationImg = document.createElement("img");
    educationImg.src = school.img;
    educationField.appendChild(educationImg);

    const educationText = document.createElement("div");
    educationField.appendChild(educationText);

    Object.keys(school)
      .filter((key) => key !== "img")
      .map((key) => {
        const child = document.createElement("div");
        child.className = `textField ${getStyleClass(key)}`;
        child.innerHTML +=
          (key == "specialization"
            ? norwegian
              ? "Spesialisering: "
              : "Specialization: "
            : "") + school[key];
        educationText.appendChild(child);
      });

    educationFields.appendChild(educationField);
  });
  data.experience.map((job) => {
    const experienceField = document.createElement("li");
    experienceField.className = "row spaced dataField";

    const experienceImg = document.createElement("img");
    experienceImg.src = job.img;
    experienceField.appendChild(experienceImg);

    const experienceText = document.createElement("div");
    experienceField.appendChild(experienceText);

    Object.keys(job)
      .filter((field) => field !== "img")
      .map((field) => {
        const jobElement = document.createElement("div");
        jobElement.className = `textField ${getStyleClass(field)}`;
        if (field == "reference") {
          const referenceContainer = document.createElement("div");
          referenceContainer.className = "row spaced";
          jobElement.appendChild(referenceContainer);

          const referenceField = document.createElement("div");
          referenceField.textContent = norwegian ? "Referanse:" : "Reference:";
          referenceField.className = "bold";
          referenceContainer.appendChild(referenceField);

          const infoContainer = document.createElement("div");
          referenceContainer.appendChild(infoContainer);

          const nameField = document.createElement("div");
          nameField.textContent = `${job[field].name} (${job[field].title})`;
          infoContainer.appendChild(nameField);

          if (job[field].phonenumber) {
            const phoneField = document.createElement("div");
            phoneField.textContent = `${norwegian ? "Tlf." : "Phone"}: ${
              job[field].phonenumber
            }`;
            infoContainer.appendChild(phoneField);
          }

          if (job[field].email) {
            const emailField = document.createElement("div");
            emailField.textContent = `${norwegian ? "E-post" : "Email"}: ${
              job[field].email
            }`;
            infoContainer.appendChild(emailField);
          }
        } else if (field == "link") {
          const text = document.createTextNode(`${job[field].text}: `);
          const link = document.createElement("a");
          link.textContent = job[field].linkText;
          link.href = job[field].url;
          jobElement.appendChild(text);
          jobElement.appendChild(link);
        } else {
          jobElement.innerHTML += job[field];
        }
        experienceText.appendChild(jobElement);
      });

    experienceFields.appendChild(experienceField);
  });
};
