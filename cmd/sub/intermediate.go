package sub

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateIntermediateFile() string {
	f, err := os.CreateTemp(os.TempDir(), "pvault_*.yaml")
	if err != nil {
		panic("temp intermediate file for the ansible format couldn't be created, abbort")
	}
	ret := f.Name()
	f.Close()
	return ret
}

type replaceState int

const (
	just_write replaceState = iota
	found_linebreak
)

const lineBreak string = "|-"
const vaultStr string = "!vault"
const ansibleLineBreak string = "!vault |-"
const ansibleVaultTxt string = "$ANSIBLE_VAULT;"

func CreateOutputFromIntermediate(input, output string) error {
	inFile, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inFile.Close()

	outFile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	state := just_write
	lastLine := ""

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, lineBreak) {
			state = found_linebreak
			lastLine = line
			continue
		}

		switch state {
		case just_write:
			if _, err := writer.WriteString(line + "\n"); err != nil {
				return fmt.Errorf("failed to write to output file: %w", err)
			}
		case found_linebreak:
			lastLineToWrite := lastLine
			if strings.Contains(line, ansibleVaultTxt) {
				// line after the line break indicates a vault
				if !strings.Contains(lastLine, vaultStr) {
					lastLineToWrite = strings.Replace(lastLine, lineBreak, ansibleLineBreak, 1)
				}
				lastLine = ""
			}
			if _, err := writer.WriteString(lastLineToWrite + "\n"); err != nil {
				return fmt.Errorf("failed to write to output file: %w", err)
			}
			if _, err := writer.WriteString(line + "\n"); err != nil {
				return fmt.Errorf("failed to write to output file: %w", err)
			}
			state = just_write
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	return nil
}
