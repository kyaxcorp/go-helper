package folder

import (
	"github.com/kyaxcorp/go-core/core/helpers/file"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"os"
)

type ScanAndCleanResult struct {
	NrOfDeletedItems   int
	NrOfUndeletedItems int
	DeletedItems       []string
	UndeletedItems     []string
}

// ScanAndClean -> will scan the folder for items, and it will check if there are multiple files in it.... if there
// are more than maxNrOfItems then it will delete them by scanned order...
func ScanAndClean(
	path string,
	maxNrOfItems int,
	//onDelete func(item fs.FileInfo, itemPath string),
	onDelete func(item os.DirEntry, itemPath string),
	//onScan func(item fs.FileInfo, itemPath string),
	onScan func(item os.DirEntry, itemPath string),
	onFinish func(result ScanAndCleanResult),
) error {
	// Scan the directory
	//files, _err := ioutil.ReadDir(path)
	files, _err := os.ReadDir(path)
	if _err != nil {
		// log.Fatal(_err)
		return _err
	}

	isOnDelete := false
	isOnScan := false
	isOnFinish := false
	if function.IsCallable(onDelete) {
		isOnDelete = true
	}
	if function.IsCallable(onScan) {
		isOnScan = true
	}
	if function.IsCallable(onFinish) {
		isOnFinish = true
	}

	// Find how many items are there!
	nrOfItems := len(files)
	nrOfItemsToDelete := 0
	// Check if it's higher than the max value
	if nrOfItems > maxNrOfItems {
		nrOfItemsToDelete = nrOfItems - maxNrOfItems
	}

	//log.Println("folders", len(files))

	// These are the items that are already being deleted
	// Here we store the items that have being deleted
	var deletedItems []string
	// Here we store the items that haven't being able to delete
	var undeletedItems []string
	// Loop through files
	for _, f := range files {
		// Create the full item path
		itemPath := file.FilterPath(path + filesystem.DirSeparator() + f.Name())
		// Call the onScan callback
		if isOnScan {
			onScan(f, itemPath)
		}

		// Check if we have items for deletion by checking the counter...
		if nrOfItemsToDelete > 0 {
			nrOfItemsToDelete = nrOfItemsToDelete - 1

			// Call the callback
			if isOnDelete {
				onDelete(f, itemPath)
			}

			//log.Println("Deleting...", itemPath)
			// Delete the item!
			if _err := file.RemoveRecursive(itemPath); _err == nil {
				deletedItems = append(deletedItems, itemPath)
			} else {
				undeletedItems = append(undeletedItems, itemPath)
			}
		}

		// fmt.Println(f.Name())

	}

	// We can create a report of results:
	// Nr of Deleted items
	// Nr Of undeleted items
	// Nr of scanned items
	//

	if isOnFinish {
		onFinish(ScanAndCleanResult{
			NrOfDeletedItems:   len(deletedItems),
			NrOfUndeletedItems: len(undeletedItems),
			DeletedItems:       deletedItems,
			UndeletedItems:     undeletedItems,
		})
	}

	return nil
}
