package main

// func TestProcess(t *testing.T) {
// 	provider := repository.Provider{
// 		Server:   "localhost",
// 		Database: "testdb",
// 	}
// 	err := provider.Connect()
// 	if err != nil {
// 		t.Errorf("Can't connect to database: %v", err)
// 	}
// 	err = provider.Truncate()
// 	if err != nil {
// 		t.Errorf("Can't truncate database")
// 	}

// 	randNumb := rand.Intn(100000)
// 	randFileName := fmt.Sprintf("/tmp/logs-%s.txt", strconv.Itoa(randNumb))
// 	fmt.Printf("Random file name: %v\n", randFileName)
// 	file, err := findOrCreateFile(randFileName)
// 	if err != nil {
// 		t.Errorf("Failed with error: %v", err)
// 	}
// 	args := []string{
// 		"program",
// 		randFileName,
// 		"1",
// 	}

// 	go startMonitoringFiles(&provider, args)

// 	expected := 40
// 	for i := 0; i < expected; i++ {
// 		// generateRandomLog(file)
// 		lg := loggenerator.LogGenerator{
// 			Writer: file,
// 		}

// 		err := lg.Run()

// 		if err != nil {
// 			t.Errorf("Failed with error: %v", err)
// 		}

// 		time.Sleep(time.Millisecond * 5)
// 	}

// 	records, err := provider.FindAll(expected+1, 0)

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if len(records) != expected {
// 		t.Errorf("Error while storing on database, expected %v, received %v", expected, len(records))
// 	}
// }
