- It looks like you wait until all results have been received from the goroutines before starting to merge them. In fact, you can (and should) start processing as soon as you receive the first result.