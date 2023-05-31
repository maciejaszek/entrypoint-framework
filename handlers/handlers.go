/*
 *  Copyright (c) 2023 Samsung Electronics Co., Ltd All Rights Reserved
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License
 */

// Package handlers provides three types of handlers.
//
// ActivationHandler provides information of a current state (active or inactive) of an application.
//
// ConfigurationHandler provides information about changes made to configuration and allows to update it in a consistent way.
// Single file ConfigurationHandler is intended for solutions where only one configuration file is present.
// Tarred ConfigurationHandler is used when configuration contains of multiple files which are provided as a tar.
// Custom ConfigurationHandler is used when a user needs to run some custom actions file while updating.
//
// ProcessHandler TODO
package handlers

import (
	"log/slog"

	"github.com/k-lb/entrypoint-framework/handlers/internal/filesystem"
	"github.com/k-lb/entrypoint-framework/handlers/internal/global"
)

const (
	handlerLogKey   = "handler"
	typeKey         = "type"
	hardlinkPostfix = "_hardlink"
)

// ActivationHandler provides information of a current state (active or inactive) of application. It can be closed with
// Close method.
type ActivationHandler interface {
	GetWasChangedChannel() <-chan ActivationEvent
	Close()
}

// ActivationEvent contains a current state of an activation (active or inactive) and an error if it was observed.
type ActivationEvent struct {
	State bool
	Error error
}

// NewActivationHandler returns a new ActivationHandler and an error if any occurred. Activation is changed based on
// presence of an activationFile.
func NewActivationHandler(activationFile string, logger *slog.Logger) (*FileActivationHandler, error) {
	log := global.HandleNilLogger(logger).With(slog.String(handlerLogKey, "activation"), slog.String("file", activationFile))
	return newFileActivationHandler(activationFile, log, filesystem.New(log))
}

// ConfigurationHandler provides methods to safely update a configuration. It should be used when the configuration is
// written and read by different application and locking mechanism can't be used (e.g. two docker containers with shared
// volume). A new configuration file should only be moved to by writer and hardlinked by reader. ConfigurationHandler
// provides information about changes made to a configuration file. It allows to update it in a consistent way and to
// get update result. It can be closed with Close method.
type ConfigurationHandler[T any] interface {
	GetWasChangedChannel() <-chan error
	Update()
	GetUpdateResultChannel() <-chan T
	Close()
}

// NewSingleFileConfigurationHandler returns a new ConfigurationHandler and an error if any occurred. Changes to
// a newConfig will be watched and when Update is called it will be copied to oldConfig which is safe to read and write
// if no update is ongoing.
func NewSingleFileConfigurationHandler(newConfig, oldConfig string, logger *slog.Logger) (*ConfigurationHandlerBase[error], error) {
	log := global.HandleNilLogger(logger).With(
		slog.String(handlerLogKey, "configuration"),
		slog.String(typeKey, "single file"),
		slog.String("newConfig", newConfig),
		slog.String("oldConfig", oldConfig))
	fs := filesystem.New(log)
	hardlink := newConfig + hardlinkPostfix
	return newConfigurationHandlerBase(
		newConfig, hardlink, updateSingleFileConfig(hardlink, oldConfig, fs), log, fs)
}

// NewTarredConfigurationHandler returns a new ConfigurationHandler and an error if any occurred. Changes to
// a newConfigFile will be watched and when Update is called it will extract newConfigFile to newConfigDir and compare
// and update its content to an oldConfigDir. newConfigDir and oldConfigDir must be on the same device.
func NewTarredConfigurationHandler(newConfigFile, newConfigDir, oldConfigDir string, logger *slog.Logger) (*ConfigurationHandlerBase[UpdateResult], error) {
	log := global.HandleNilLogger(logger).With(
		slog.String(handlerLogKey, "configuration"),
		slog.String(typeKey, "tarred"),
		slog.String("newConfigFile", newConfigFile),
		slog.String("newConfigDir", newConfigDir),
		slog.String("oldConfigDir", oldConfigDir))
	fs := filesystem.New(log)
	hardlink := newConfigFile + hardlinkPostfix
	return newConfigurationHandlerBase(
		newConfigFile, hardlink, updateTarredConfig(hardlink, newConfigDir, oldConfigDir, fs), log, fs)
}

// NewCustomConfigurationHandler returns a new ConfigurationHandler and an error if any occurred. Changes to
// a newConfigFile will be watched and a hardlink will be created of this file. The update function will be called by
// ConfigurationHandler.Update().
func NewCustomConfigurationHandler[T any](newConfigFile, hardlink string, update func() T, logger *slog.Logger) (*ConfigurationHandlerBase[T], error) {
	log := global.HandleNilLogger(logger).With(
		slog.String(handlerLogKey, "configuration"),
		slog.String(typeKey, "custom"),
		slog.String("newConfigFile", newConfigFile),
		slog.String("hardlink", hardlink))
	return newConfigurationHandlerBase(
		newConfigFile, hardlink, update, log, filesystem.New(log))
}
