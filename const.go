package main

const (
	DEBUG = false
	DEBUG_SAMPLE_JSON = `{"PurchaseOrder":{"-OrderDate":"1999-10-20","-PurchaseOrderNumber":"99503","Address":[{"-Type":"Shipdasdasping","City":"MillValley","Country":"USA","Name":"EllenAdams","State":"CA","Street":"123MapleStreet","Zip":"10999"},{"-Type":"Billing","City":"OldTown","Country":"USA","Name":"TaiYee","State":"PA","Street":"8OakAvenue","Zip":"95819"}],"DeliveryNotes":"Pleaseleavepackagesinshedbydriveway.","Items":{"Item":[{"-PartNumber":"872-AA","Comment":"Confirmthisiselectric","ProductName":"Lawnmower","Quantity":"1","USPrice":"148.95"},{"-PartNumber":"926-AA","ProductName":"BabyMonitor","Quantity":"2","ShipDate":"1999-05-21","USPrice":"39.98"}]}}}`
	DEBUG_SAMPLE_XML = `<?xml version="1.0"?><PurchaseOrder PurchaseOrderNumber="99503" OrderDate="1999-10-20">    <Address Type="Shipping">        <Name>Ellen Adams</Name>        <Street>123 Maple Street</Street>        <City>Mill Valley</City>        <State>CA</State>        <Zip>10999</Zip>        <Country>USA</Country>    </Address>    <Address Type="Billing">        <Name>Tai Yee</Name>        <Street>8 Oak Avenue</Street>        <City>Old Town</City>        <State>PA</State>        <Zip>95819</Zip>        <Country>USA</Country>    </Address>    <DeliveryNotes>Please leave packages in shed by driveway.</DeliveryNotes>    <Items>        <Item PartNumber="872-AA">            <ProductName>Lawnmower</ProductName>            <Quantity>1</Quantity>            <USPrice>148.95</USPrice>            <Comment>Confirm this is electric</Comment>        </Item>        <Item PartNumber="926-AA">            <ProductName>Baby Monitor</ProductName>            <Quantity>2</Quantity>            <USPrice>39.98</USPrice>            <ShipDate>1999-05-21</ShipDate>        </Item></Items></PurchaseOrder>`

	APP_NAME      = "jsonConverter"
	MAJOR_VERSION = "1"
	MINOR_VERSION = "0"
 	VERSION       = "v" + MAJOR_VERSION + "." + MINOR_VERSION

 	INTI_WINDOW_HEIGHT = 700.0
 	INTI_WINDOW_WIDTH = 700.0

 	FONT_ENV_NAME        = "FYNE_FONT"
 	FONT_WIN_MEIRYO_PATH = `C:\Windows\Fonts\meiryo.ttc`

 	PREFIX_DEFAULT = ""
 	INDENT_DEFAULT = "  "
)
