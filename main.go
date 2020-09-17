package main

import "os"

/*
* creates a folder with infocard title
 */
func createFolder(list []string) {
	//create folder for each item in the list (document number)

	//add subfolder Attachments

}

/*
* changes file name from [infocardid].[ext] to [infocardnum].[ext]
 */
func changeFileName() {

}

/*
* adds main file (original and pdf) to folder
 */
func addMainFile() {

}

/*
* add attachments to attachments subfolder
 */
func addAttachments() {

}

/*
* put earlier revisions into folder of new revisions
 */
func addRevisions() {

}

func main() {

	//create EFP folder
	os.Mkdir("EFP", 0755)

	//create folder for each infocard - revision

}
