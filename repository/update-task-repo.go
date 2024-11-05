package repository

import "fmt"

func (r *TaskRepositoryDB) UpdateStatusTask(id int) error {
	// Query untuk memperbarui status task menjadi 'complete' berdasarkan ID
	query := "UPDATE tasks SET status = 'complete' WHERE id = $1"

	// Menjalankan query update
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error updating task status: %v", err)
	}

	// Jika tidak ada error, return nil untuk menandakan bahwa update berhasil
	return nil
}
