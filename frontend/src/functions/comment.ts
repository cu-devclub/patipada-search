type CommentPosition = {
  index: number;
  position: number;
};

export const getStartAndEndIndexOfComments = (html: string) => {
  // loop over character in string
  const temp: CommentPosition[][] = [];
  let cnt = 0;
  for (let i = 0; i < html.length; i++) {
    // if p tag; ignore
    if (html[i] == "<" && html[i + 1] == "p") {
      i += 2;
      continue;
    } else if (html[i] == "<" && html[i + 1] == "/" && html[i + 2] == "p") {
      i += 3;
      continue;
    }

    // if span tag ; move to >
    else if (html[i] == "<" && html[i + 1] == "s") {
      while (html[i] != ">") {
        i++;
      }
      i += 1;
      const startPos: CommentPosition = { index: i, position: cnt + 1 };

      while (html[i] != "<") {
        i++;
        cnt++;
      }
      const endPos: CommentPosition = { index: i, position: cnt };
      temp.push([startPos, endPos]);
    } else {
      // increment cnt for normal character
      cnt++;
    }
  }

  return temp;
};

export const removeCommentFromHTML = (
  html: string,
  start: number,
  end: number,
  time: string
) => {
  let i = start - 1;
  while (html[i] != "<") {
    i--;
  }
  let j = end;
  while (html[j] != ">") {
    j++;
  }

  // finding the list of comments
  let m = i;
  for (let k = i; k < start; k++) {
    if (html[k] == "[") {
      m = k;
      break;
    }
  }
  // split list , if there is only one comment we delete the whole span
  // else we delete only the comment that have match with time
  const q = [m + 1];
  for (let n = m + 1; n < html.length; n++) {
    if (html[n] == "}" && html[n + 1] == "," && html[n + 2] == "{") {
      q.push(n + 1);
      break;
    }
  }

  if (q.length == 1) {
    // only one; delete the whole span
    const newHTMLString =
      html.slice(0, i) + html.slice(start, end) + html.slice(j + 1);
    return newHTMLString;
  }

  // loop over q to find the time
  q.push(start - 3);
  for (let s = 0; s < q.length; s++) {
    const tex = html.slice(q[s], q[s + 1]);
    if (tex.includes(time)) {
      // Delete the comment
      const frontText = html.slice(0, q[s]);
      const backText = html.slice(q[s + 1]);
      if (frontText[frontText.length - 1] == ",") {
        return html.slice(0, q[s]-1) + html.slice(q[s + 1])
      } 
      if (backText[0] == ",") {
        return html.slice(0, q[s]) + html.slice(q[s+1]+1)
      }
      else {
        return frontText + backText;
      }
    } 
  }

  return html;
};