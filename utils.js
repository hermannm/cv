const addElement = ({
  type,
  parent,
  id,
  className,
  textContent,
  href,
  src,
  width,
  height,
}) => {
  const element = document.createElement(type);
  if (parent) parent.appendChild(element);
  if (id) element.id = id;
  if (className) element.className = className;
  if (textContent) element.textContent = textContent;
  if (href) element.href = href;
  if (src) element.src = src;
  if (width) element.width = width;
  if (height) element.height = height;
  return element;
};

const addText = ({ parent, paragraph }) => {
  if (Array.isArray(paragraph)) {
    for (const subField of paragraph) {
      if (typeof subField === "object") {
        if (subField.link) {
          addElement({
            type: "a",
            parent: parent,
            textContent: subField.text,
            href: subField.link,
          });
        } else if (subField.bold) {
          addElement({
            type: "strong",
            parent: parent,
            textContent: subField.text,
          });
        } else if (subField.italics) {
          addElement({
            type: "em",
            parent: parent,
            textContent: subField.text,
          });
        }
      } else {
        parent.appendChild(document.createTextNode(subField));
      }
    }
  } else {
    parent.textContent = paragraph;
  }
};

const addIconElement = ({ iconKey, iconColor, textContent, link, parent }) => {
  const container = addElement({
    type: link ? "a" : "div",
    parent,
    className: "row spaced verticallyCentered",
    href: link,
  });

  let iconPath = undefined;
  switch (iconKey) {
    case "age":
      iconPath = "icons/about.svg";
      break;
    case "email":
      iconPath = `icons/email_${iconColor}.svg`;
      break;
    case "phone":
      iconPath = `icons/phone_${iconColor}.svg`;
      break;
    case "github":
      iconPath = "icons/github.svg";
      break;
    case "linkedin":
      iconPath = "icons/linkedin.svg";
      break;
  }
  if (iconPath) {
    addElement({
      type: "img",
      parent: container,
      src: iconPath,
      width: "20",
      height: "20",
    });
  }

  addElement({
    type: "div",
    parent: container,
    textContent,
  });

  return container;
};

const addListSection = ({ title, list, mainContainer, mapFunctionCreator }) => {
  addElement({
    type: "div",
    parent: mainContainer,
    className: "sectionTitle",
    textContent: title,
  });

  const fields = addElement({
    type: "ul",
    parent: mainContainer,
    className: "column fieldGap",
  });

  list.map((listItem) => {
    const field = addElement({
      type: "li",
      parent: fields,
      className: "row fieldGap",
    });

    addElement({
      type: "img",
      parent: field,
      className: "logo",
      src: listItem.img,
    });

    const text = addElement({
      type: "div",
      parent: field,
    });

    const mapFunction = mapFunctionCreator(text);

    Object.entries(listItem)
      .filter(([key]) => key !== "img")
      .map(mapFunction);
  });
};

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
