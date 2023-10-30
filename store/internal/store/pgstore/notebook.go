package pgstore

import (
	"database/sql"
	log "store/internal/logger"
	"store/internal/models"
)

type NotebookRepo struct {
	db *sql.DB
}

func (r *NotebookRepo) CreateNotebook(n *models.Notebook) error {
	if err := n.Validate(); err != nil {
		return err
	}

	_, err := r.db.Exec(
		`insert into notebooks 
			(systems_id, screens_id, processors_id, storages_id, rams_id, model, description, price)
		values
			($1, $2, $3, $4, $5, $6, $7, $8)`,
		n.SystemID, n.ScreenID, n.ProcessorID, n.StorageID, n.RAMID, n.Model, n.Description, n.Price)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	return nil
}

func (r *NotebookRepo) DeleteNotebookByID(id string) error {
	_, err := r.db.Exec("delete from notebooks where id_notebooks = $1", id)
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	return nil
}

func (r *NotebookRepo) FindNotebookByID(id string, info *models.NotebookInfo) error {
	err := r.db.QueryRow(
		`select 
			s.name, sc.size_inches, sc.resolution, p.model, p.speed_ghz, 
			st.type_storage, st.size_gb, r.size_gb, n.model, n.description, n.price
		 from 
			notebooks as n
		 inner join systems as s on n.systems_id = s.id_systems
		 inner join screens as sc on n.screens_id = sc.id_screens
		 inner join processors as p on n.processors_id = p.id_processors
		 inner join storages as st on n.storages_id = st.id_storages
		 inner join rams as r on n.rams_id = r.id_rams
		 where 
			n.id_notebooks = $1;`, id).Scan(
		&info.SystemName, &info.ScreenSizeInches, &info.ScreenResolution,
		&info.ProcessorModel, &info.ProcessorSpeedGHz, &info.StorageType,
		&info.StorageSizeGB, &info.RAMSizeGB, &info.NotebookModel,
		&info.NotebookDescription, &info.NotebookPrice)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	return nil
}

func (r *NotebookRepo) GetAllNotebooks() ([]*models.NotebookInfo, error) {
	var notebooks []*models.NotebookInfo
	rows, err := r.db.Query(
		`select 
            s.name, sc.size_inches, sc.resolution, p.model, p.speed_ghz, 
            st.type_storage, st.size_gb, r.size_gb, n.model, n.description, n.price
         from 
            notebooks as n
         inner join systems as s on n.systems_id = s.id_systems
         inner join screens as sc on n.screens_id = sc.id_screens
         inner join processors as p on n.processors_id = p.id_processors
         inner join storages as st on n.storages_id = st.id_storages
         inner join rams as r on n.rams_id = r.id_rams`)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		info := &models.NotebookInfo{}
		err := rows.Scan(
			&info.SystemName, &info.ScreenSizeInches, &info.ScreenResolution,
			&info.ProcessorModel, &info.ProcessorSpeedGHz, &info.StorageType,
			&info.StorageSizeGB, &info.RAMSizeGB, &info.NotebookModel,
			&info.NotebookDescription, &info.NotebookPrice)
		if err != nil {
			log.Logger.Error(err)
			return nil, err
		}
		notebooks = append(notebooks, info)
	}

	return notebooks, nil
}
