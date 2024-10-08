
// Code generated by mtgroup-generator.
package dal

import (
	"time"
		"database/sql"
		"demo/internal/types"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

const (

	sqlGetExample = `
	SELECT
		id,
		cpu_units,
		memory_mb,
		name,
		price,
		storage_mb,
		user_id
	FROM
		examples
	WHERE
		id=:id AND
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted
	`

	sqlGetMyExampleID = `
	SELECT
		id
	FROM
		examples
	WHERE
		created_by=:created_by AND
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted AND
		bound
	`


	sqlAddExample = `
	INSERT INTO examples(
		cpu_units,
		memory_mb,
		name,
		price,
		storage_mb,
		user_id,
		id,
		created_by,
		isolated_entity_id
	) VALUES (
		:cpu_units,
		:memory_mb,
		:name,
		:price,
		:storage_mb,
		:user_id,
		:id,
		:created_by,
		:isolated_entity_id
	)
	RETURNING
		id
	`

	sqlBindExampleToProfile = `
	UPDATE
		examples
	SET
		bound=true
	WHERE
		id=:id AND
		created_by=:created_by AND
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted AND
		NOT bound
	`

	sqlDeleteExample = `
	UPDATE
		examples
	SET
		deleted=true,
		deleted_at=:deleted_at,
		deleted_by=:deleted_by
	WHERE
		id=:id AND
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted
	`

	sqlEditExample = `
	UPDATE
		examples
	SET
		cpu_units=:cpu_units,
		memory_mb=:memory_mb,
		name=:name,
		price=:price,
		storage_mb=:storage_mb,
		user_id=:user_id
	WHERE
		id=:id AND
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted
	`
	
	sqlListExample = `
	SELECT
		id,
		cpu_units,
		memory_mb,
		name,
		price,
		storage_mb,
		user_id
	FROM
		examples
	WHERE
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted
	`
	sqlListExampleCount = `
	SELECT
		COUNT(*)
	FROM
		examples
	WHERE
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted
	`
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
			
					sqlGetUserForExampleLazyLoading = `
	SELECT
		id,
		balance,
		name,
		role
	FROM
		users
	WHERE
		id=:id AND
		isolated_entity_id=:isolated_entity_id
	`
			
	
)

type (
    argGetExample struct {
        ID sql.NullString `db:"id"`
			IsolatedEntityID string `db:"isolated_entity_id"`
    }

	argGetMyExampleID struct {
			CreatedBy string `db:"created_by"`
			IsolatedEntityID string `db:"isolated_entity_id"`
    }
	argAddExample struct {
							ID string `db:"id"`
							CpuUnits types.Decimal `db:"cpu_units"`
							MemoryMb int32 `db:"memory_mb"`
							Name string `db:"name"`
							Price types.Decimal `db:"price"`
							StorageMb int32 `db:"storage_mb"`
					UserID interface{} `db:"user_id"`
			CreatedBy string `db:"created_by"`
			IsolatedEntityID string `db:"isolated_entity_id"`
	}
	argBindExampleToProfile struct {
		ID string `db:"id"`
			CreatedBy string `db:"created_by"`
			IsolatedEntityID string `db:"isolated_entity_id"`
	}
	argEditExample struct {
						ID string `db:"id"`
						CpuUnits types.Decimal `db:"cpu_units"`
						MemoryMb int32 `db:"memory_mb"`
						Name string `db:"name"`
						Price types.Decimal `db:"price"`
						StorageMb int32 `db:"storage_mb"`
					UserID interface{} `db:"user_id"`
			IsolatedEntityID string `db:"isolated_entity_id"`
	}
	argDeleteExample struct {
		ID string `db:"id"`
		DeletedAt *time.Time `db:"deleted_at"`
		DeletedBy string `db:"deleted_by"`
			IsolatedEntityID string `db:"isolated_entity_id"`
	}
)