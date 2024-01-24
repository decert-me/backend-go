package service

import (
	"fmt"
	"github.com/adrg/sysfont"
	"image/color"
	"log"
	"regexp"
	"strings"
	"testing"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

func Test_Test(t *testing.T) {
	finder := sysfont.NewFinder(nil)

	for _, font := range finder.List() {
		fmt.Println(font.Family, font.Name, font.Filename)
	}
}

// 1024-(textPad*2)
func Test_generateImage(t *testing.T) {
	//// 打开图片文件
	//file, err := os.Open("/Users/mac/Desktop/bafkreid4lhm7bpv3o7ycfk55b64mkl5ahbxjgf6bdvvphk2i4becg7ms3u.png")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//
	//// 解码图片
	//img, _, err := image.Decode(file)
	//if err != nil {
	//	log.Fatal(err)
	//}

	width := 1024
	height := 1024

	// 创建画布
	dc := gg.NewContext(width, height)
	//dc := gg.NewContext(width, height)
	//dc.Scale(float64(width)/float64(img.Bounds().Dx()), float64(height)/float64(img.Bounds().Dy()))
	img, err := gg.LoadImage("/Users/mac/Desktop/bafkreid4lhm7bpv3o7ycfk55b64mkl5ahbxjgf6bdvvphk2i4becg7ms3u.png")
	dstImage128 := imaging.Resize(img, width, height, imaging.Lanczos)
	if err := dc.LoadFontFace("msyh", 64); err != nil {
		fmt.Println("Error loading font:", err)
		return
	}
	dc.DrawImage(dstImage128, 0, 0)

	// 定义渐变遮罩层的高度和颜色
	startColor := color.RGBA{0, 0, 0, 255} // 渐变起始颜色
	endColor := color.RGBA{0, 0, 0, 0}     // 渐变结束颜色

	// 创建线性渐变画刷
	linearGradient := gg.NewLinearGradient(0, float64(height), 0, float64(height)-float64(height)*0.23)
	linearGradient.AddColorStop(0, startColor)
	linearGradient.AddColorStop(1, endColor)

	// 绘制渐变遮罩层
	dc.SetFillStyle(linearGradient)
	fmt.Println(0, float64(height)-maskHeight, float64(width), maskHeight)
	dc.DrawRectangle(0, float64(height)-maskHeight, float64(width), maskHeight)
	dc.Fill()

	// 绘制文本
	text := "编写 ERC721 NFT 合约"
	dc.SetRGB(1, 1, 1)
	splitWords(dc, text)
	dc.Fill()
	// 保存为 PNG 图片
	err = dc.SavePNG("gradient.png")
	if err != nil {
		log.Fatal("保存图片失败:", err)
	}

	log.Println("渐变背景图已创建")
}

// dc.SavePNG("out.png")
// 将图像保存到文件
// saveImageToFile(img, "output.png")
func splitWords(ctx *gg.Context, text string) {
	// 分割
	re := regexp.MustCompile(`\p{Han}\s?|^\p{Han}\s?|\S+\s?`)
	words := re.FindAllString(text, -1)
	var line string
	var lines []string

	for i := 0; i < len(words); i++ {
		testLine := line + words[i]
		testWidth, _ := ctx.MeasureString(testLine)
		if testWidth > textMaxWidth && i > 0 {
			lines = append(lines, strings.TrimSpace(line))
			line = words[i]
			if len(lines) == 2 {
				// 如果已经有两行，将省略号添加到第二行的末尾并终止循环
				lines[1] = strings.TrimSpace(lines[1]) + "..."
				break
			}
		} else {
			line = testLine
		}
	}
	if len(lines) < 2 {
		lines = append(lines, strings.TrimSpace(line))
	}
	for i := 0; i < len(lines); i++ {
		lineHeight := textHeight
		if len(lines) == 1 {
			lineHeight = textHeight * 2
		}
		textpd := (float64(lineHeight)-64.0)/2.0 + (20.0 * float64(i))
		ctx.DrawStringAnchored(lines[i], textPad, (px*0.88)+(float64(i)*float64(lineHeight))+textpd, 0, 0)
	}
}
