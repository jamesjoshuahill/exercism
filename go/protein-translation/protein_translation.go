package protein

const (
	// ErrStop indicates a STOP codon.
	ErrStop = err("stop")
	// ErrInvalidBase indicates a codon without a corresponding amino acid.
	ErrInvalidBase = err("invalid base")
)

type err string

func (e err) Error() string { return string(e) }

// FromCodon returns the amino acid corresponding to the codon, or an error:
// ErrStop if it is a STOP codon, or ErrInvalidBase if it is invalid.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA returns a list of amino acids corresponding to the codons in a RNA
// sequence, or an error. In the case of an error, the amino acids identified
// before the error is returned.
func FromRNA(rna string) ([]string, error) {
	p := make([]string, len(rna)/3)
	for i := 0; i < len(rna); i += 3 {
		base := rna[i : i+3]
		c, err := FromCodon(base)
		if err != nil {
			p = p[:i/3]
			if err == ErrStop {
				break
			}
			return p, err
		}
		p[i/3] = c
	}
	return p, nil
}
