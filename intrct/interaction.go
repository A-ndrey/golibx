package intrct

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/A-ndrey/golibx/slicex"
)

type Interaction struct {
	Stdin           io.Reader
	Stdout          io.Writer
	UserInputPrefix string
}

func (i Interaction) PromptBool(question string, def bool) (bool, error) {
	opts := "[y/N]"
	if def {
		opts = "[Y/n]"
	}

	if _, err := fmt.Fprintf(i.Stdout, "%s %s: ", question, opts); err != nil {
		return false, fmt.Errorf("writing question: %w", err)
	}

	userInput, err := bufio.NewReader(i.Stdin).ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return def, nil
		}
		return false, fmt.Errorf("reading user input: %w", err)
	}

	userInput = strings.ToLower(strings.TrimSpace(userInput))

	if userInput == "" {
		return def, nil
	}

	if userInput == "y" || userInput == "yes" {
		return true, nil
	}

	return false, nil
}

func (i Interaction) PromptList(question string, opts []string, def []string) ([]string, error) {
	const chooseOptionsHint = "Choose option(s) (e.g. 1 2 4)"
	orderedOpts := slicex.Map(opts, func(i int, s string) string {
		return fmt.Sprintf("%d. %s", i+1, s)
	})

	userAnswer, err := i.PromptString(fmt.Sprintf("%s:\n%s\n%s", question, strings.Join(orderedOpts, "\n  "), chooseOptionsHint), "")
	if err != nil {
		return nil, err
	}

	if userAnswer == "" {
		return def, nil
	}

	indices, err := parseOptionsIndices(userAnswer)
	if err != nil {
		return nil, err
	}

	var res []string
	for _, i := range indices {
		if i < 0 || i >= len(opts) {
			return nil, fmt.Errorf("invalid option index: %d", i)
		}
		res = append(res, opts[i])
	}

	return res, nil
}

func parseOptionsIndices(opts string) ([]int, error) {
	var res []int
	optFields := strings.Fields(opts)
	for _, optField := range optFields {
		num, err := strconv.ParseInt(optField, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, int(num-1))
	}

	slices.Sort(res)
	slices.Compact(res)

	return res, nil
}

func (i Interaction) PromptString(question string, def string) (string, error) {
	if _, err := fmt.Fprintln(i.Stdout, question); err != nil {
		return "", fmt.Errorf("writing question: %w", err)
	}

	if _, err := fmt.Fprint(i.Stdout, i.UserInputPrefix); err != nil {
		return "", fmt.Errorf("writing prefix: %w", err)
	}

	userInput, err := bufio.NewReader(i.Stdin).ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return def, nil
		}
		return "", fmt.Errorf("reading user input: %w", err)
	}

	userInput = strings.TrimSpace(userInput)

	if userInput == "" {
		return def, nil
	}

	return userInput, nil
}

var Default = Interaction{
	Stdin:           os.Stdin,
	Stdout:          os.Stdout,
	UserInputPrefix: ">>> ",
}

func PromptBool(question string, def bool) (bool, error) {
	return Default.PromptBool(question, def)
}

func PromptList(question string, opts []string, def []string) ([]string, error) {
	return Default.PromptList(question, opts, def)
}

func PromptString(question string, def string) (string, error) {
	return Default.PromptString(question, def)
}
