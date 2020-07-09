// Identify authentication and authorization service
//
// Copyright (C) 2020 Alexei Broner
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package config

import (
	"fmt"
	"os"
	"path"
)

func GetHTTPAddress() string {
	address := os.Getenv("IDENTIFY_HTTP_ADDRESS")
	if len(address) == 0 {
		return ":8080"
	}
	return address
}

func GetRealm() string {
	realm := os.Getenv("IDENTIFY_REALM")
	if len(realm) == 0 {
		return "localhost"
	}
	return realm
}

func GetTokenSecret() (string, error) {
	tokenSecret := os.Getenv("IDENTIFY_TOKEN_SECRET")
	if len(tokenSecret) == 0 {
		return "", fmt.Errorf("An secret key to sign tokens with must be " +
			"provided by the environment variable IDENTIFY_TOKEN_SECRET.")
	}
	return tokenSecret, nil
}

func GetCredentialsPath() (string, error) {
	credsPath := os.Getenv("IDENTIFY_CREDENTIALS_PATH")
	if len(credsPath) == 0 {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("A path to a credentials file must be provided" +
				"by the environment variable IDENTIFY_CREDENTIALS_PATH.")
		}
		return path.Join(home, ".identify", "credentials.json"), nil
	}
	return credsPath, nil
}

func GetDBPath() (string, error) {
	dbPath := os.Getenv("IDENTIFY_DB_PATH")
	if len(dbPath) == 0 {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("A path to an identity database file must be " +
				"provided by the environment variable IDENTIFY_DB_PATH.")
		}
		return path.Join(home, ".identify", "identity.db"), nil
	}
	return dbPath, nil
}

func GetTokenDBPath() (string, error) {
	tokenDBPath := os.Getenv("IDENTIFY_TOKEN_DB_PATH")
	if len(tokenDBPath) == 0 {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("A path to an identity database file must be " +
				"provided by the environment variable IDENTIFY_TOKEN_DB_PATH.")
		}
		return path.Join(home, ".identify", "token.db"), nil
	}
	return tokenDBPath, nil
}