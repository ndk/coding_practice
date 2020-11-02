package exercise

func numUniqueEmails(emails []string) int {
	normalized := map[string]struct{}{}
	for _, email := range emails {
		normalizedEmail := make([]byte, len(email))
		end := 0
		for i := 0; i < len(email); i++ {
			if email[i] == '@' {
				for ; i < len(email); i++ {
					normalizedEmail[end] = email[i]
					end++
				}
				normalized[string(normalizedEmail[:end])] = struct{}{}
				break
			}
			if email[i] == '.' {
				continue
			}
			if email[i] == '+' {
				for i++; i < len(email) && email[i] != '@'; i++ {
				}
				i--
				continue
			}
			normalizedEmail[end] = email[i]
			end++
		}
	}
	return len(normalized)
}
