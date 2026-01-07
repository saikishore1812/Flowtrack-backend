package utils

import (
	"flowtrack-backend/config"
	"flowtrack-backend/models"
)

func LogAudit(userID uint, actorName, entity string, entityID uint) {
	log := models.AuditLog{
		UserID:   userID,
		
		ActorName: actorName,
		Entity:  entity,
		EntityID: entityID,
	}
	config.DB.Create(&log)
}
