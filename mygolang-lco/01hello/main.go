package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

//func main() {
//	offerLabelPath := "dunzo@@icons@@home@@dunzodaily@@offer@Star3_%s_%.0f.png"
//
//	offerLabelPath = fmt.Sprintf(offerLabelPath, "flat", 5.0)
//
//	newTransformation := fmt.Sprintf("h-$bh$,w-$bw$,cm-pad_resize:oi-%s,ox-40,oy-40,oh-440", offerLabelPath)
//	var re = regexp.MustCompile(`tr:(.*?)/`)
//	fmt.Println(re)
//
//	// {{tr1}} is used to for replacement, because $ in newTransformation is getting considered as expand operator during regex replace
//	imageUrl := "https://ik.imagekit.io/dunzo/tr:w-$w$,h-$h$,cm-pad_resize/dGg0NE0rdWlGZjBnZlJjNmE4eWQ2QT09-product_image.jpg"
//	// h-$bh$,w-$bw$,cm-pad_resize:oi-dunzo@@icons@@home@@dunzodaily@@offer@Star3_${offerType}_${offer['value']}.png,ox-40,oy-40,oh-440
//	imageUrl = re.ReplaceAllString(imageUrl, fmt.Sprintf("tr:{{tr1}}:$1/"))
//	fmt.Println(imageUrl)
//	imageUrl = strings.Replace(imageUrl, "{{tr1}}", newTransformation, 1)
//	fmt.Println(imageUrl)
//
//	'https://ik.imagekit.io/dunzo/dunzo-web/public/junk/smallbrandcardnoshadow4x.png/tr:h-288,w-288,r-32:oi-dunzo-web@@public@@junk@@OfferTag2.png,ofo-top_left:ot-10%25%20OFF,otf-Gilroy-Bold.ttf,ots-40,otc-FFFFFF,ofo-top_left,otp-0_16_0_16:w-200,h-144'
//	'https://ik.imagekit.io/dunzo/tr:h-$bh$,w-$bw$,cm-pad_resize:oi-dunzo@@icons@@home@@dunzodaily@@offer@Star3_flat_5.png,ox-40,oy-40,oh-440:w-$w$,h-$h$,cm-pad_resize/dGg0NE0rdWlGZjBnZlJjNmE4eWQ2QT09-product_image.jpg
//
//
////`https://ik.imagekit.io/dunzo/dunzo-web/tr:h-$bh$,w-$bw$,cm-pad_resize:oi-dunzo@@icons@@home@@dunzodaily@@offer@Star3_percent_5.png,ox-40,oy-40,oh-440:w-$w$,h-$h$,cm-pad_resize/public/junk/categoryImage.png`
////'          https://ik.imagekit.io/dunzo/tr:h-$bh$,w-$bw$,cm-pad_resize:oi-dunzo@@icons@@home@@dunzodaily@@offer@Star3_flat_5.png,ox-40:w-$w$,h-$h$,cm-pad_resize/dGg0NE0rdWlGZjBnZlJjNmE4eWQ2QT09-product_image.jpg
////'

const (
	UptoEndIndex                   = 6
	startIndex                     = 0
	OfferLabelTextWithoutUpto      = "%v%s OFF"
	OfferLabelTextWithUPTO         = "UP TO %v%s OFF"
	OfferLabelTextWithUptoNewline  = "UP TO\n %v%s OFF"
	OfferLabelStandardTextWithUpto = "Upto %v%s OFF"
	OfferLabelStandardTextNewline  = "%v%s\n OFF"
	OfferLabelTransformation       = "/tr:h-$bh$,w-$bw$,ot-%v,otf-Gilroy-Bold.ttf,ots-30,otc-FFFFFF,otbg-412A84,otp-18_32_18_32,bg-FFFFFF,ox-$rx$,oy-432,or-20,ofo-center:w-$w$,h-$h$"
	OfferLabelOldTransformation    = "/tr:w-$w$,h-$h$_home_icon"
	LabelFontName                  = "gilroy-semibold"
	LabelUptoTextFontSize          = 8
	LabelTextFontSize              = 10
	PercentageSign                 = "%"
)

func createImageUrlWithOfferLabel(imageUrl string, offer string, Value float64) *string {
	offerLabelPath := "dunzo@@icons@@home@@dunzodaily@@offer@Star3_%s_%.0f.png"
	if offer == "FLAT" {
		offerLabelPath = fmt.Sprintf(offerLabelPath, "flat", Value)
	} else {
		offerLabelPath = fmt.Sprintf(offerLabelPath, "percent", Value)
	}
	newTransformation := fmt.Sprintf("h-$bh$,w-$bw$,cm-pad_resize:oi-%s,ox-40,oy-40,oh-440", offerLabelPath)
	var re = regexp.MustCompile(`tr:(.*?)/`)

	// {{tr1}} is used to for replacement, because $ in newTransformation is getting considered as expand operator during regex replace
	imageUrl = re.ReplaceAllString(imageUrl, fmt.Sprintf("tr:{{tr1}}:$1/"))
	imageUrl = strings.Replace(imageUrl, "{{tr1}}", newTransformation, 1)
	fmt.Println(imageUrl)
	return &imageUrl
}

func test() {
	// You can edit this code!

	value := 14.0
	imageUrl := "https://ik.imagekit.io/dunzo/f01118852ac411edb0eb80894534762c_PRODUCT_631073e8fb73ba00139689f6_1.JPG"
	offerLabelPath := "dunzo@@icons@@home@@dunzodaily@@offer@Star3_%s_%.0f.png"

	offerLabelPath = fmt.Sprintf(offerLabelPath, "percent", value)

	newTransformation := fmt.Sprintf("h-$bh$,w-$bw$,cm-pad_resize:oi-%s,ox-40,oy-40,oh-440", offerLabelPath)
	var re = regexp.MustCompile(`tr:(.*?)/`)

	// {{tr1}} is used to for replacement, because $ in newTransformation is getting considered as expand operator during regex replace
	imageUrl = re.ReplaceAllString(imageUrl, fmt.Sprintf("tr:{{tr1}}:$1/"))
	imageUrl = strings.Replace(imageUrl, "{{tr1}}", newTransformation, 1)
	fmt.Println(imageUrl)

}

func test2() {
	value := 14.0
	//imageUrl := "https://ik.imagekit.io/dunzo/f01118852ac411edb0eb80894534762c_PRODUCT_631073e8fb73ba00139689f6_1.JPG"
	imageUrl := "https://ik.imagekit.io/dunzo/tr:w-$w$,h-$h$,cm-pad_resize/1657640425343_product_62cd67bacdf550002fd13692_1.jpg -"
	offerLabelPath := "dunzo@@icons@@home@@dunzodaily@@offer@Star3_%s_%.0f.png"

	offerLabelPath = fmt.Sprintf(offerLabelPath, "percent", value)

	newTransformation := fmt.Sprintf("h-$bh$,w-$bw$,cm-pad_resize:oi-%s,ox-40,oy-40,oh-440", offerLabelPath)
	trLocation := regexp.MustCompile("tr:(.)*\\/").FindStringIndex(imageUrl)

	//overwriting existing transformation
	newImageUrl := fmt.Sprint(imageUrl[:trLocation[0]-1] + newTransformation + imageUrl[trLocation[1]-1:])
	fmt.Println(newImageUrl)
}

func main() {
	//val := "https://ik.imagekit.io/dunzo/dunzo-new-stage/tr:w-$w$,h-$h$_home_icon/25285ec8-012e-4d76-a75f-d537420d7fbe.png"
	//val1 := CreateRoundedImageKitWithBgColor(val, "", "")
	//fmt.Println(val1)
	//test := "/tr:w-$w$,h-$h$_home_icon"
	//val2 := strings.ReplaceAll(val, test, "")
	//val2 := "UP TO\n"
	//fmt.Println(len(val2))

	//var val []string
	//val = append(val, "jake")
	//val = append(val, "paul")
	//val = append(val, "logan")
	//val = append(val, "baul")
	//var val2 []string
	//
	//val2 = append(val2, "ok")
	//for _, i := range val {
	//	val2 = append(val2, i)
	//}
	test2()

	//fmt.Println(val2)

}

func newImageUrl() {
	existingItemImageUrl := "https://ik.imagekit.io/dunzo/dunzo-new-stage/tr:w-$w$,h-$h$_home_icon/25285ec8-012e-4d76-a75f-d537420d7fbe.png"
	trLocation := ImageKitTransformRegex.FindStringIndex(existingItemImageUrl)
	offerPercentage := 15
	offerLabelText := fmt.Sprintf(OfferLabelTextWithUPTO, offerPercentage, PercentageSign)
	t := &url.URL{Path: offerLabelText}
	urlEncodedOfferText := t.String()

	newTransformation := fmt.Sprintf(OfferLabelTransformation, urlEncodedOfferText)
	//overwriting existing transformation
	newImageUrl := fmt.Sprint(existingItemImageUrl[:trLocation[0]-1] + newTransformation + existingItemImageUrl[trLocation[1]-1:])
	fmt.Println(newImageUrl)
}

var (
	ImageKitTransformRegex = regexp.MustCompile("tr:(.)*\\/")
)

func CreateRoundedImageKitWithBgColor(imageKitUrl string, bgColor string, cornerRadius string) string {
	bgColor = strings.TrimPrefix(bgColor, "#")
	trLocation := ImageKitTransformRegex.FindStringIndex(imageKitUrl)
	fmt.Println(trLocation, trLocation[1])
	fmt.Println(imageKitUrl[:trLocation[1]-1])
	fmt.Println(imageKitUrl[trLocation[1]-1:])
	if trLocation != nil {
		return fmt.Sprintf("%v:bg-%v,r-%v%v", imageKitUrl[:trLocation[1]-1], strings.ToUpper(bgColor), cornerRadius,
			imageKitUrl[trLocation[1]-1:])
	}
	return imageKitUrl
}

func offetest() {
	//const baseUrl = "https://ik.imagekit.io/dunzo/home/icons/home/dunzodaily/dd_night_store_shut_banner.png"
	//const imagekitUrlSuffix = "tr=ote-%s,otbg-EBF7F2,otp-5,ots-50,otc-05563C,or-30,ox-450,oy-300,ott-b,otf-home@@Gilroy-Medium.ttf:w-$w$,h-$h$_home_banner"

	//DefaultStoreCloseTime := "10PM"
	//DefaultStoreOpenTime := "6AM"
	//timingString := fmt.Sprintf("%s - %s", DefaultStoreOpenTime, DefaultStoreCloseTime)
	//fmt.Println(timingString)
	//sEnc := base64.StdEncoding.EncodeToString([]byte(timingString))
	//
	//fmt.Println(sEnc)
	//fmt.Println(baseUrl)
	//fmt.Println(imagekitUrlSuffix)
	//imageUrl := fmt.Sprintf(baseUrl+"?"+imagekitUrlSuffix, sEnc)
	//
	//fmt.Println(imageUrl)
	////
	//fmt.Println()
	offerPercentage := 100
	//OfferLabelTextWithUPTO := `UP%20TO%s20%v%s%s20OFF`
	existingItemImageUrl := "https://ik.imagekit.io/dunzo/dunzo-new-stage/25285ec8-012e-4d76-a75f-d537420d7fbe.png"
	testing := fmt.Sprintf("UP TO %v%s OFF", offerPercentage, "%")
	//offerText := fmt.Sprintf(testing, "%", "%", offerPercentage, "%", "%")
	//testing := `UPTO%2010%25%20OFF`
	println(testing)
	sEnc := base64.StdEncoding.EncodeToString([]byte(testing))
	fmt.Println(sEnc)

	test := url.QueryEscape(testing)
	fmt.Println(test)

	params := url.Values{}
	params.Add("message", testing)

	//test1, _ := UrlEncoded(testing)
	//fmt.Println(test1)

	temp, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(temp)

	mySecondString := testing
	t := &url.URL{Path: mySecondString}
	mySecondEncodedString := t.String()

	fmt.Println(mySecondEncodedString)

	//newTransformation := "/tr:h-752,w-544,ot-" + mySecondEncodedString + ",otf-Gilroy-Bold.ttf,ots-30,otc-FFFFFF,otbg-412A84,otp-18_32_18_32,bg-FFFFFF,ox-124,oy-432,or-20,ofo-center:w-272,h-392"
	//
	////part1 := "/tr:h-752,w-544,ote-"
	////part2 := `,otf-Gilroy-Bold.ttf,ots-30,otc-FFFFFF,otbg-412A84,otp-18_32_18_32,bg-FFFFFF,ox-124,oy-432,or-20,ofo-center:w-272,h-392\`
	////newTransformation := fmt.Sprintf("%v?%v", part1, sEnc, part2)
	//newImageUrl := existingItemImageUrl + newTransformation
	//fmt.Println(newImageUrl)

	OfferLabelTransformation := "/tr:h-$bh$,w-$bw$,ot-%v,otf-Gilroy-Bold.ttf,ots-30,otc-FFFFFF,otbg-412A84,otp-18_32_18_32,bg-FFFFFF,ox-$rx$,oy-432,or-20,ofo-center:w-$w$,h-$h$"
	filledOfferLabelTransformation := "/tr:h-752,w-544,ot-%s,otf-Gilroy-Bold.ttf,ots-30,otc-FFFFFF,otbg-412A84,otp-18_32_18_32,bg-FFFFFF,ox-124,oy-432,or-20,ofo-center:w-272,h-392"
	fmt.Println(filledOfferLabelTransformation)

	newTransformation := fmt.Sprintf(OfferLabelTransformation, mySecondEncodedString)
	fmt.Println(newTransformation)
	//part1 := "/tr:h-752,w-544,ote-"
	//part2 := `,otf-Gilroy-Bold.ttf,ots-30,otc-FFFFFF,otbg-412A84,otp-18_32_18_32,bg-FFFFFF,ox-124,oy-432,or-20,ofo-center:w-272,h-392\`
	//newTransformation := fmt.Sprintf("%v?%v", part1, sEnc, part2)
	newImageUrl := fmt.Sprint(existingItemImageUrl + newTransformation)
	fmt.Println(newImageUrl)

}
