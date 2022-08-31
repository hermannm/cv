import * as data from "../data.js";
import { addElement, addIconElement, addListSection, addText, getStyleClass } from "./utils.js";

const sidebar = document.getElementById("sidebar");
const mainContainer = document.getElementById("main");

addElement({ type: "img", parent: sidebar, src: data.person.profilepicture });

const personFields = addElement({ type: "div", parent: sidebar, className: "padded column fieldGap" });

for (const [key, item] of Object.entries(data.person)) {
  switch (key) {
    case "profilepicture":
    case "signature":
      break;
    case "name":
      addElement({
        type: "div",
        parent: personFields,
        className: `textField ${getStyleClass(key)}`,
        textContent: item.text,
      });
      break;
    default:
      addIconElement({
        iconKey: key,
        iconColor: "white",
        textContent: item.text,
        link: item.link,
        parent: personFields,
      });
  }
}

if (data.application) {
  addElement({
    type: "div",
    parent: mainContainer,
    textContent: `${data.english ? "Application" : "Søknad"}: ${data.application}`,
    className: "sectionTitle",
  });

  for (const applicationParagraph of data.applications[data.application]) {
    const paragraphElement = addElement({ type: "p", parent: mainContainer });
    addText({ parent: paragraphElement, paragraph: applicationParagraph });
  }

  addElement({ type: "div", parent: mainContainer, textContent: data.english ? "Sincerely," : "Med vennlig hilsen," });

  addElement({ type: "img", parent: mainContainer, src: data.person.signature, id: "signature" });

  addElement({ type: "div", parent: mainContainer, textContent: data.person.name.text });
} else {
  addListSection({
    title: data.english ? "Education" : "Utdanning",
    list: data.education,
    mainContainer,
    listItemTransformer: (textParent) => (key, item) => {
      addElement({
        type: "div",
        parent: textParent,
        className: `textField ${getStyleClass(key)}`,
        textContent: `
            ${key === "specialization" ? (data.english ? "Specialization: " : "Spesialisering: ") : ""}${item}`,
      });
    },
  });

  addListSection({
    title: data.english ? "Experience" : "Erfaring",
    list: data.experience,
    mainContainer,
    listItemTransformer: (textParent) => (key, field) => {
      const experienceTextItem = addElement({
        type: "div",
        parent: textParent,
        className: `textField ${getStyleClass(key)}`,
      });

      if (key === "reference") {
        const referenceContainer = addElement({ type: "div", parent: experienceTextItem, className: "row spaced" });

        addElement({
          type: "div",
          parent: referenceContainer,
          className: "bold",
          textContent: data.english ? "Reference:" : "Referanse:",
        });

        const infoContainer = addElement({ type: "div", parent: referenceContainer });

        addElement({ type: "div", parent: infoContainer, textContent: `${field.name} (${field.title})` });

        if (field.phone || field.email) {
          const contactContainer = addElement({ type: "div", parent: infoContainer, className: "row fieldGap" });

          if (field.phone) {
            addIconElement({
              iconKey: "phone",
              iconColor: "black",
              parent: contactContainer,
              textContent: field.phone,
            });
          }

          if (field.email) {
            addIconElement({
              iconKey: "email",
              iconColor: "black",
              parent: contactContainer,
              textContent: field.email,
              link: `mailto:${field.email}`,
            });
          }
        }
      } else {
        addText({ parent: experienceTextItem, paragraph: field });
      }
    },
  });
}
