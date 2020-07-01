//
// This is only a SKELETON file for the 'RNA Transcription' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

const complements = {
  C: "G",
  G: "C",
  T: "A",
  A: "U",
};

export const toRna = (dna) => {
  let rna = "";
  for (const nucleotide of dna) {
    rna += complements[nucleotide];
  }
  return rna;
};
