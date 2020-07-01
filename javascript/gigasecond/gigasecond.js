//
// This is only a SKELETON file for the 'Gigasecond' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

const trillion = 1e9;

export const gigasecond = (moment) => {
  const newTime = moment.getTime() + trillion * 1000
  return new Date(newTime)
};
