package v14

/*

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import (
	"github.com/apache/trafficcontrol/lib/go-log"
	"testing"
)

func TestUsersDeliveryServices(t *testing.T) {
	CreateTestCDNs(t)
	CreateTestTypes(t)
	CreateTestProfiles(t)
	CreateTestStatuses(t)
	CreateTestDivisions(t)
	CreateTestRegions(t)
	CreateTestPhysLocations(t)
	CreateTestCacheGroups(t)
	CreateTestDeliveryServices(t)

	CreateTestUsersDeliveryServices(t)
	GetTestUsersDeliveryServices(t)
	DeleteTestUsersDeliveryServices(t)

	DeleteTestDeliveryServices(t)
	DeleteTestCacheGroups(t)
	DeleteTestPhysLocations(t)
	DeleteTestRegions(t)
	DeleteTestDivisions(t)
	DeleteTestStatuses(t)
	DeleteTestProfiles(t)
	DeleteTestTypes(t)
	DeleteTestCDNs(t)
}

const TestUsersDeliveryServicesUser = "admin" // TODO make dynamic

func CreateTestUsersDeliveryServices(t *testing.T) {
	log.Debugln("CreateTestUsersDeliveryServices")

	dses, _, err := TOSession.GetDeliveryServices()
	if err != nil {
		t.Fatalf("cannot GET DeliveryServices: %v - %v\n", err, dses)
	}
	if len(dses) == 0 {
		t.Fatalf("no delivery services, must have at least 1 ds to test users_deliveryservices\n")
	}
	users, _, err := TOSession.GetUsers()
	if err != nil {
		t.Fatalf("cannot GET users: %v\n", err)
	}
	if len(users) == 0 {
		t.Fatalf("no users, must have at least 1 user to test users_deliveryservices\n")
	}

	dsIDs := []int{}
	for _, ds := range dses {
		dsIDs = append(dsIDs, ds.ID)
	}

	userID := 0
	foundUser := false
	for _, user := range users {
		if *user.Username == TestUsersDeliveryServicesUser {
			userID = *user.ID
			foundUser = true
			break
		}
	}
	if !foundUser {
		t.Fatalf("get users expected: %v actual: missing\n", TestUsersDeliveryServicesUser)
	}

	_, err = TOSession.SetDeliveryServiceUser(userID, dsIDs, true)
	if err != nil {
		t.Fatalf("failed to set delivery service users: " + err.Error())
	}

	userDSes, _, err := TOSession.GetUserDeliveryServices(userID)
	if err != nil {
		t.Fatalf("get user delivery services returned error: " + err.Error())
	}

	if len(userDSes.Response) != len(dsIDs) {
		t.Fatalf("get user delivery services expected %v actual %v\n", len(dsIDs), len(userDSes.Response))
	}

	actualDSIDMap := map[int]struct{}{}
	for _, userDS := range userDSes.Response {
		if userDS.ID == nil {
			t.Fatalf("get user delivery services returned a DS with a nil ID\n")
		}
		actualDSIDMap[*userDS.ID] = struct{}{}
	}
	for _, dsID := range dsIDs {
		if _, ok := actualDSIDMap[dsID]; !ok {
			t.Fatalf("get user delivery services expected %v actual %v\n", dsID, "missing")
		}
	}
}

func GetTestUsersDeliveryServices(t *testing.T) {
	log.Debugln("GetTestUsersDeliveryServices")

	dses, _, err := TOSession.GetDeliveryServices()
	if err != nil {
		t.Fatalf("cannot GET DeliveryServices: %v - %v\n", err, dses)
	}
	if len(dses) == 0 {
		t.Fatalf("no delivery services, must have at least 1 ds to test users_deliveryservices\n")
	}
	users, _, err := TOSession.GetUsers()
	if err != nil {
		t.Fatalf("cannot GET users: %v\n", err)
	}
	if len(users) == 0 {
		t.Fatalf("no users, must have at least 1 user to test users_deliveryservices\n")
	}

	dsIDs := []int64{}
	for _, ds := range dses {
		dsIDs = append(dsIDs, int64(ds.ID))
	}

	userID := 0
	foundUser := false
	for _, user := range users {
		if *user.Username == TestUsersDeliveryServicesUser {
			userID = *user.ID
			foundUser = true
			break
		}
	}
	if !foundUser {
		t.Fatalf("get users expected: %v actual: missing\n", TestUsersDeliveryServicesUser)
	}

	userDSes, _, err := TOSession.GetUserDeliveryServices(userID)
	if err != nil {
		t.Fatalf("get user delivery services returned error: " + err.Error() + "\n")
	}

	if len(userDSes.Response) != len(dsIDs) {
		t.Fatalf("get user delivery services expected %v actual %v\n", len(dsIDs), len(userDSes.Response))
	}

	actualDSIDMap := map[int]struct{}{}
	for _, userDS := range userDSes.Response {
		if userDS.ID == nil {
			t.Fatalf("get user delivery services returned a DS with a nil ID\n")
		}
		actualDSIDMap[*userDS.ID] = struct{}{}
	}
	for _, dsID := range dsIDs {
		if _, ok := actualDSIDMap[int(dsID)]; !ok {
			t.Fatalf("get user delivery services expected %v actual %v\n", dsID, "missing")
		}
	}
}

func DeleteTestUsersDeliveryServices(t *testing.T) {
	log.Debugln("DeleteTestUsersDeliveryServices")

	users, _, err := TOSession.GetUsers()
	if err != nil {
		t.Fatalf("cannot GET users: %v\n", err)
	}
	if len(users) == 0 {
		t.Fatalf("no users, must have at least 1 user to test users_deliveryservices\n")
	}
	userID := 0
	foundUser := false
	for _, user := range users {
		if *user.Username == TestUsersDeliveryServicesUser {
			userID = *user.ID
			foundUser = true
			break
		}
	}
	if !foundUser {
		t.Fatalf("get users expected: %v actual: missing\n", TestUsersDeliveryServicesUser)
	}

	dses, _, err := TOSession.GetUserDeliveryServices(userID)
	if err != nil {
		t.Fatalf("get user delivery services returned error: " + err.Error())
	}
	if len(dses.Response) == 0 {
		t.Fatalf("get user delivery services expected %v actual %v\n", ">0", "0")
	}

	for _, ds := range dses.Response {
		if ds.ID == nil {
			t.Fatalf("get user delivery services returned ds with nil ID\n")
		}
		_, err := TOSession.DeleteDeliveryServiceUser(userID, *ds.ID)
		if err != nil {
			t.Fatalf("delete user delivery service returned error: " + err.Error())
		}
	}

	dses, _, err = TOSession.GetUserDeliveryServices(userID)
	if err != nil {
		t.Fatalf("get user delivery services returned error: " + err.Error())
	}
	if len(dses.Response) != 0 {
		t.Fatalf("get user delivery services after deleting expected %v actual %v\n", "0", len(dses.Response))
	}
}
