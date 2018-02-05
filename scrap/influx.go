package main

import (
	"time"
	 "github.com/influxdata/influxdb/client/v2"
)

const (
	DbName = "db1"
	TableName = "apts"
)

type dbInfo struct  {
	dbName string
	tableName string
	client  client.Client
	batchPoints client.BatchPoints
}


func DbInit() (*dbInfo, error) {
        // Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",

	})
	if err != nil {
		return nil, err
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  DbName,
		Precision: "s",
	})
	if err != nil {
		return nil, err
	}

	return &dbInfo{dbName:DbName, tableName: TableName, client:c, batchPoints:bp}, nil
}

func (db *dbInfo)Write(pts []point) error {

	for _, p := range pts {
		pt, err := client.NewPoint(TableName, p.tags, p.fields, time.Now())
		if err != nil {
			return err
		}
		db.batchPoints.AddPoint(pt)
	}

	// Write the batch
	return db.client.Write(db.batchPoints)
}
