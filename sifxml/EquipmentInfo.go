package sifxml


    type EquipmentInfo struct {
        RefId RefIdType `xml:"RefId,attr"`
      Name string `xml:"Name"`
      Description string `xml:"Description"`
      LocalId LocalIdType `xml:"LocalId"`
      AssetNumber LocalIdType `xml:"AssetNumber"`
      InvoiceRefId string `xml:"InvoiceRefId"`
      PurchaseOrderRefId string `xml:"PurchaseOrderRefId"`
      EquipmentType string `xml:"EquipmentType"`
      SIF_RefId EquipmentInfo_SIF_RefId `xml:"SIF_RefId"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    type EquipmentInfo_SIF_RefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
