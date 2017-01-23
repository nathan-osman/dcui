package docker

import (
	"sort"

	"golang.org/x/net/context"
)

// ServiceStatus represents the status of a particular service.
type ServiceStatus struct {
	Name  string
	State string
}

type byName []*ServiceStatus

func (b byName) Len() int           { return len(b) }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byName) Less(i, j int) bool { return b[i].Name < b[j].Name }

// Status retrieves the status of all services.
func (d *Docker) Status() ([]*ServiceStatus, error) {
	infos, err := d.project.Ps(context.Background())
	if err != nil {
		return nil, err
	}
	stateMap := make(map[string]string)
	for _, i := range infos {
		if n, ok := i["Name"]; ok {
			s, _ := i["State"]
			stateMap[n] = s
		}
	}
	statusList := []*ServiceStatus{}
	for k, v := range d.project.ServiceConfigs.All() {
		s, _ := stateMap[v.ContainerName]
		statusList = append(statusList, &ServiceStatus{
			Name:  k,
			State: s,
		})
	}
	sort.Sort(byName(statusList))
	return statusList, nil
}
