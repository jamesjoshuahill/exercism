//
// This is only a SKELETON file for the 'Protein Translation' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const translate = (bases) => {
  if (!bases) { return [] }

  const codons = bases.match(/.../g)
  const anticodons = codons.map((c) => anticodon(c))

  return trimAfterStop(anticodons)
};

const STOP = 'STOP';

const anticodon = (c) => {
  switch (c) {
    case 'AUG':
      return 'Methionine'
    case 'UUU':
    case 'UUC':
      return 'Phenylalanine'
    case 'UUA':
    case 'UUG':
      return 'Leucine'
    case 'UCU':
    case 'UCC':
    case 'UCA':
    case 'UCG':
      return 'Serine'
    case 'UAU':
    case 'UAC':
      return 'Tyrosine'
    case 'UGU':
    case 'UGC':
      return 'Cysteine'
    case 'UGG':
      return 'Tryptophan'
    case 'UAA':
    case 'UAG':
    case 'UGA':
      return STOP
    default:
      throw new Error('Invalid codon')
  }
}

const trimAfterStop = (anticodons) => {
  const stopIndex = anticodons.indexOf(STOP)
  if (stopIndex != -1) {
    return anticodons.slice(0, stopIndex)
  }

  return anticodons
}
