const addElement = ({
  type,
  parent,
  className,
  textContent,
  href,
  src,
  id,
}) => {
  const element = document.createElement(type);
  if (parent) parent.appendChild(element);
  if (id) element.id = id;
  if (className) element.className = className;
  if (textContent) element.textContent = textContent;
  if (href) element.href = href;
  if (src) element.src = src;
  return element;
};

const addText = ({ parent, paragraph }) => {
  if (Array.isArray(paragraph)) {
    for (const subField of paragraph) {
      if (typeof subField === "object") {
        if (subField.url) {
          addElement({
            type: "a",
            parent: parent,
            textContent: subField.text,
            href: subField.url,
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

const addIconElement = ({ iconPath, text, link, parent }) => {
  const container = addElement({
    type: link ? "a" : "div",
    parent,
    className: "row spaced",
    href: link,
  });

  addElement({
    type: "img",
    parent: container,
    src: iconPath,
  });

  addElement({
    type: "div",
    parent: container,
    textContent: text,
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
