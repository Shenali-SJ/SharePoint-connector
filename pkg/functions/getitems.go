package functions

import (
	"fmt"
	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip/api"
	"sharepoint-connector/pkg/env"
	"sharepoint-connector/pkg/model"
)

func GetItems(listname string) ([]*model.Listitemdata, error) {

	client := env.Initreturnvalue.(*gosip.SPClient)

	sp := api.NewSP(client)
	list := sp.Web().Lists().GetByTitle(listname)  // gets list by list name
	itemsResp, err := list.Items().Select("Id, Title").Get()  // gets Id and Title of each item in the list

	if err != nil {
		fmt.Println(err)
	}

	var listItemData []*model.Listitemdata

	for _, item := range itemsResp.Data() {  // unmarshalls each item and gets Id and Title
		itemsData := item.Data()
		itemModel := &model.Listitemdata{
			Title:  itemsData.Title,
			Itemid: int64(itemsData.ID),
		}

		listItemData := append(listItemData, itemModel)
		fmt.Println(listItemData)
	}

	return listItemData, nil
}
