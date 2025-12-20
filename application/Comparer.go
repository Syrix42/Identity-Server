package application


type Comparer interface {
    Compare(hashed string, plaintext string) error
}