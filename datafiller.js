window.onload = () => {
  const sidebar = document.getElementById("sidebar");
  const mainContainer = document.getElementById("main");

  addElement({
    type: "img",
    parent: sidebar,
    src: data.person.profilepicture,
  });

  const personFields = addElement({
    type: "div",
    parent: sidebar,
    className: "padded column fieldGap",
  });

  Object.entries(data.person)
    .filter(([key]) => !(key == "profilepicture" || key == "signature"))
    .map(([key, item]) => {
      if (key == "name") {
        addElement({
          type: "div",
          parent: personFields,
          className: `textField ${getStyleClass(key)}`,
          textContent: item.text,
        });
      } else {
        addIconElement({
          iconKey: key,
          iconColor: "white",
          text: item.text,
          link: item.link,
          parent: personFields,
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
      mapFunctionCreator:
        (textParent) =>
        ([key, item]) => {
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
            } ${item}`,
          });
        },
    });

    addListSection({
      title: english ? "Experience" : "Erfaring",
      list: data.experience,
      mainContainer,
      mapFunctionCreator:
        (textParent) =>
        ([key, item]) => {
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
              textContent: `${item.name} (${item.title})`,
            });

            if (item.phone) {
              addIconElement({
                iconKey: "phone",
                iconColor: "black",
                parent: infoContainer,
                text: item.phone,
              });
            }

            if (item.email) {
              addIconElement({
                iconKey: "email",
                iconColor: "black",
                parent: infoContainer,
                text: item.email,
                link: `mailto:${item.email}`,
              });
            }
          } else {
            addText({ parent: experienceTextItem, paragraph: item });
          }
        },
    });
  }
};
