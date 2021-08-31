window.onload = () => {
  const mainContainer = document.getElementById("main");

  const profilePicture = document.getElementById("profilepicture");
  profilePicture.src = data.person.profilepicture;

  const personFields = document.getElementById("person");

  Object.keys(data.person)
    .filter((field) => !(field == "profilepicture" || field == "signature"))
    .map((field) => {
      const personField = addElement({
        type: "li",
        parent: personFields,
        className: `textField ${getStyleClass(field)}`,
        textContent: data.person[field].text,
      });
      if (data.person[field].link) {
        addElement({
          type: "a",
          parent: personField,
          href: data.person[field].link,
        });
      }
    });

  if (application) {
    addElement({
      type: "div",
      parent: mainContainer,
      textContent: `${english ? "Application" : "Søknad"}: ${application}`,
      className: "sectionTitle",
    });

    const applicationText = addElement({
      type: "div",
      parent: mainContainer,
    });

    for (const applicationParagraph of applications[application]) {
      const paragraphElement = addElement({
        type: "p",
        parent: applicationText,
      });
      addText({ parent: paragraphElement, paragraph: applicationParagraph });
    }

    addElement({
      type: "div",
      parent: applicationText,
      textContent: english ? "Sincerely," : "Med vennlig hilsen,",
    });

    addElement({
      type: "img",
      parent: applicationText,
      src: data.person.signature,
      id: "signature",
    });

    addElement({
      type: "div",
      parent: applicationText,
      textContent: data.person.name.text,
    });
  } else {
    addListSection({
      title: english ? "Education" : "Utdanning",
      list: data.education,
      mainContainer,
      mapFunctionCreator: (textParent, listItem) => (key) => {
        addElement({
          type: "div",
          parent: textParent,
          className: `textField ${getStyleClass(key)}`,
          textContent: `
            ${
              key == "specialization"
                ? english
                  ? "Specialization: "
                  : "Spesialisering: "
                : ""
            } ${listItem[key]}`,
        });
      },
    });

    addListSection({
      title: english ? "Experience" : "Erfaring",
      list: data.experience,
      mainContainer,
      mapFunctionCreator: (textParent, listItem) => (key) => {
        const experienceTextItem = addElement({
          type: "div",
          parent: textParent,
          className: `textField ${getStyleClass(key)}`,
        });

        if (key == "reference") {
          const referenceContainer = addElement({
            type: "div",
            parent: experienceTextItem,
            className: "row spaced",
          });

          addElement({
            type: "div",
            parent: referenceContainer,
            className: "bold",
            textContent: english ? "Reference:" : "Referanse:",
          });

          const infoContainer = addElement({
            type: "div",
            parent: referenceContainer,
          });

          addElement({
            type: "div",
            parent: infoContainer,
            textContent: `${listItem[key].name} (${listItem[key].title})`,
          });

          if (listItem[key].phonenumber) {
            addElement({
              type: "div",
              parent: infoContainer,
              textContent: `${english ? "Phone" : "Tlf."}: ${
                listItem[key].phonenumber
              }`,
            });
          }

          if (listItem[key].email) {
            const emailField = addElement({
              type: "div",
              parent: infoContainer,
            });
            addText({
              parent: emailField,
              paragraph: [
                `${english ? "Email" : "E-post"}: `,
                {
                  text: listItem[key].email,
                  url: `mailto:${listItem[key].email}`,
                },
              ],
            });
          }
        } else if (key == "link") {
          const text = document.createTextNode(`${listItem[key].text}: `);
          const link = document.createElement("a");
          link.textContent = listItem[key].linkText;
          link.href = listItem[key].url;
          experienceTextItem.appendChild(text);
          experienceTextItem.appendChild(link);
        } else {
          experienceTextItem.innerHTML += listItem[key];
        }
      },
    });
  }
};
