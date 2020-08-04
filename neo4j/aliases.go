/*
 * Copyright (c) 2002-2020 "Neo4j,"
 * Neo4j Sweden AB [http://neo4j.com]
 *
 * This file is part of Neo4j.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/connection"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

// Aliases for simpler usage (fewer imports)
// A separate type package is needed to avoid circular package references and to avoid
// unnecessary copying/conversions between structs.
type (
	Point2D       = dbtype.Point2D
	Point3D       = dbtype.Point3D
	Date          = dbtype.Date
	LocalTime     = dbtype.LocalTime
	LocalDateTime = dbtype.LocalDateTime
	Time          = dbtype.Time // AKA OffsetTime
	Duration      = dbtype.Duration
	Node          = dbtype.Node
	Relationship  = dbtype.Relationship
	Path          = dbtype.Path
	Record        = connection.Record
)