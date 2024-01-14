export const extractStringFromHTML = (html: string): string => {
  return html.replace(/<[^>]*>?/gm, "");
};

export const encodeHTMLText = (html: string): string => {
  return String(html).replace(/[&<>"']/g, (match) => {
    switch (match) {
      case "&":
        return "&amp;";
      case "<":
        return "&lt;";
      case ">":
        return "&gt;";
      case '"':
        return "&quot;";
      case "'":
        return "&#39;";
      default:
        return match;
    }
  });
}

export const decodeHTMLText = (encodedString: string): string => {
  const doc = new DOMParser().parseFromString(encodedString, "text/html");
  return doc.body.textContent || "";
};

export const convertStatusWord = (status: string): string => {
  switch (status) {
    case "pending":
      return "รอการตรวจสอบ";
    case "reviewed":
      return "ตรวจสอบเรียบร้อย";
    default:
      return "";
  }
}

