import { KeyboardEvent } from "react";

export function handleEnterKeyPress<T = Element>(f: () => void) {
  return handleKeyPress<T>(f, "Enter");
}

function handleKeyPress<T = Element>(f: () => void, key: string) {
  return (e: KeyboardEvent<T>) => {
    if (e.key === key) {
      f();
    }
  };
}
