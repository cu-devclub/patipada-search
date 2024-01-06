export const  timeToSeconds = (time) => {
  const [hours, minutes, seconds] = time.split(":").map(Number);
  return hours * 3600 + minutes * 60 + seconds;
}

export const setTimeout = (time, callback) => {
  const milliseconds = time * 1000;

  // Use the global setTimeout function
  window.setTimeout(callback, milliseconds);
};