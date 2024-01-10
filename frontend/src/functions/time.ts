export const  timeToSeconds = (time) => {
  const [hours, minutes, seconds] = time.split(":").map(Number);
  return hours * 3600 + minutes * 60 + seconds;
}

export const setTimeout = (time, callback) => {
  const milliseconds = time * 1000;

  // Use the global setTimeout function
  window.setTimeout(callback, milliseconds);
};

// create a function that receives time in this format: 00:00:00 and return 3 values: hours, minutes, seconds
export const splitTime = (time:string) => {
  const [hours, minutes, seconds] = time.split(":").map(Number);
  return { hours, minutes, seconds };
};

export const generateTime = (hours:number, minutes:number, seconds:number): string => {
  const paddedHour = hours.toString().padStart(2, "0");
  const paddedMinute = minutes.toString().padStart(2, "0");
  const paddedSecond = seconds.toString().padStart(2, "0");
  const fullTimeText = `${paddedHour}:${paddedMinute}:${paddedSecond}`;
  return fullTimeText;
}