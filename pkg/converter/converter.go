package converter

import (
	"fmt"
	"strings"

	"github.com/ak1ra24/tn2inethenge/pkg/inethenge"
	"github.com/ak1ra24/tn2inethenge/pkg/tn"
)

func Tn2inethenge(srcFile string) (inethengeJson inethenge.InetHenge, err error) {
	tnconfig, err := tn.ReadTnConfig(srcFile)
	if err != nil {
		return inethengeJson, err
	}

	linksmap := map[string]string{}

	for _, node := range tnconfig.Nodes {
		nodeName := node.Name

		var inethengeNode inethenge.Node
		inethengeNode.Name = nodeName
		inethengeNode.Icon = "https://inet-henge.herokuapp.com/images/router.png"
		inethengeJson.Nodes = append(inethengeJson.Nodes, inethengeNode)

		for _, inf := range node.Interfaces {
			source := nodeName
			infArgs := strings.Split(inf.Args, "#")
			target := infArgs[0]
			link := fmt.Sprintf("%s:%s", source, target)
			sourceInf := inf.Name
			targetInf := infArgs[1]
			infLinkInfo := fmt.Sprintf("%s:%s", sourceInf, targetInf)
			reverseLink := fmt.Sprintf("%s:%s", target, source)

			if _, ok := linksmap[reverseLink]; !ok {
				linksmap[link] = infLinkInfo
			}
		}
	}

	// links := keys(linksmap)

	for link, linkInfo := range linksmap {
		lllink := strings.Split(link, ":")
		source := lllink[0]
		target := lllink[1]

		lllinkInfo := strings.Split(linkInfo, ":")
		sourceInf := lllinkInfo[0]
		targetInf := lllinkInfo[1]

		var inethengeLink inethenge.Link

		inethengeLink.Source = source
		inethengeLink.Target = target

		ihLinkMetaInf := inethenge.Interface{Source: sourceInf, Target: targetInf}
		ihLinkMetaData := inethenge.MetaData{Interface: ihLinkMetaInf}
		inethengeLink.Meta = ihLinkMetaData
		inethengeJson.Links = append(inethengeJson.Links, inethengeLink)
	}
	// jsonBytes, err := json.MarshalIndent(inethengeJson, "", "   ")
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(jsonBytes))

	return inethengeJson, nil
}

func keys(m map[string]string) []string {
	ks := []string{}
	for k, _ := range m {
		ks = append(ks, k)
	}
	return ks
}
