package syncer

import "errors"

type SyncJob struct {
	pathsToProcess map[string]string
}

func (s *SyncJob) AddPath(path string) {
	s.pathsToProcess[path] = path
}

func (s *SyncJob) GetPaths() []string {
	var paths []string
	for key, _ := range s.pathsToProcess {
		paths = append(paths, key)
	}

	return paths
}

func (s *SyncJob) Length() int {
	return len(s.pathsToProcess)
}

func (s *SyncJob) GetFirst() (string, error) {
	for path := range s.pathsToProcess {
		return path, nil
	}

	return "", errors.New("there are no paths")
}

func (s *SyncJob) Clear() {
	s.pathsToProcess = map[string]string{}
}
