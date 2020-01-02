//
// This is only a SKELETON file for the 'All Your Base' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

const INPUT_BASE_ERROR = new Error('Wrong input base')
const OUTPUT_BASE_ERROR = new Error('Wrong output base')
const INPUT_ERROR = new Error('Input has wrong format')

export const convert = (input, inputBase, outputBase) => {
  validateBases(inputBase, outputBase)
  validateInput(input, inputBase)

  const decimal = convertToDecimal(input, inputBase)

  return convertToBase(decimal, outputBase)
};

const convertToDecimal = (input, base) => {
  return input.reduce((sum, n, index) => {
    const exponent = input.length - 1 - index
    return sum + (n * Math.pow(base, exponent))
  }, 0)
}

const convertToBase = (number, base) => {
  if (number == 0) { return [0] }

  let output = []
  while (number > 0) {
    const digit = number % base
    output.unshift(digit)

    const quotient = Math.floor(number / base)
    number = quotient
  }

  return output
}

const validateBases = (inputBase, outputBase) => {
  if (isInvalidBase(inputBase)) {
    throw INPUT_BASE_ERROR
  } else if (isInvalidBase(outputBase)) {
    throw OUTPUT_BASE_ERROR
  }
}

const isInvalidBase = (base) => {
  return !base || base < 2 || !Number.isInteger(base)
}

const validateInput = (input, base) => {
  if (input.length == 0) {
    throw INPUT_ERROR
  }

  if (input.length > 1 && input[0] == 0) {
    throw INPUT_ERROR
  }

  if (input.some((n) => n < 0 || n >= base)) {
    throw INPUT_ERROR
  }
}
