package go_task

/*https://alm.dataprev.gov.br/ccm/rpt/repository/workitem?fields=workitem/workItem[id=233537]/
  (id|summary|
  type/name|
  href|
  description|
  owner/userId|
  resolver/userId|
  creator/userId|
  category/name|
  reportableUrl|
  plannedEndDate|
  foundIn/name|timeSpent|duration|plannedStartDate|activationDate|reportableUrl|creationDate|teamArea/name|comments/content|timeSheetEntries)*/
type Task struct {
	Id string `xml:"id"`
	Href string `xml:"href"`
	Description string `xml:"description"`
	CreationDate string `xml:"creation_date"`
}

type TimeSheetEntry struct {




}
