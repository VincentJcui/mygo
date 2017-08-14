package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func cleanUrls(l string, urls []string) []string {
	var resurls []string
	/*
		清洗url的形式

		http://xxx.com/a.jpg
		//xx.com/a.jpgs
		/ststic/a.jpg
		a.jpg
	*/
	s_u, err := url.Parse(l)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range urls {
		u, err := url.Parse(v)
		if err != nil {
			log.Fatal(err)
		}
		//判断是否有http
		if u.Scheme != "" {
			resurls = append(resurls, v)
			continue
		}
		//判读是否有host
		if u.Host != "" {
			resurls = append(resurls, s_u.Scheme+":"+v)
			continue
		}
		//判断path
		if u.Path != "" {
			tmp_s_url := strings.SplitAfter(s_u.Path, "/")
			tmp_d_url := strings.SplitAfter(u.Path, "/")
			for i := 0; i < len(tmp_d_url)-1; i++ {
				if tmp_s_url[i] != tmp_d_url[i] && tmp_d_url[0] == "/" {
					resurls = append(resurls, s_u.Scheme+"://"+s_u.Host+strings.Join(tmp_s_url[:i], "")+strings.Join(tmp_d_url[i:], ""))
					break
				}
				if tmp_d_url[0] != "/" && tmp_s_url[i+1] != tmp_d_url[i] {
					resurls = append(resurls, s_u.Scheme+"://"+s_u.Host+strings.Join(tmp_s_url[:i+1+1], "")+strings.Join(tmp_d_url[i:], ""))
					break
				}
			}
			continue
		}
	}

	return resurls
}

func fetch(l string) ([]string, error) {
	var urls []string
	resp, err := http.Get(l)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("http error code:%s", resp.Status)
		return nil, err
	}
	//io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := (s.Attr("src"))
		if ok {
			urls = append(urls, link)
		} else {
			fmt.Println("src not found")
		}
	})
	return cleanUrls(l, urls), err
	//return urls, err
}

func downloadImags(urls []string, dir string) error {
	for _, i := range urls {
		//下面是获取文件名名字的方法
		//filename := strings.Split(i,"/")[len(strings.Split(i,"/"))-1]   //获取文件名
		//fmt.Println(path.Base(i) )   //获取文件名
		res, err := http.Get(i)

		if err != nil {
			log.Panic(err)
		}
		defer res.Body.Close()               //注意关闭
		if res.StatusCode != http.StatusOK { //http返回码检查
			return errors.New(res.Status)
		}
		//
		fullname := filepath.Join(dir, path.Base(i)) //直接指定生产的目录和文件名
		//os.Chdir(dir)   //切换到指定目录
		f, err := os.Create(fullname)
		if err != nil {
			log.Panic(err)
		}
		defer f.Close() //注意关闭
		io.Copy(f, res.Body)
	}
	return nil

}

func maketar(dir string, w io.Writer) error {
	basedir := filepath.Base(dir)
	compress := gzip.NewWriter(w)
	defer compress.Close()
	tr := tar.NewWriter(compress)
	defer tr.Close()
	filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		/*
			写入的tar的Fileheader
			以读取的方式打开文件
			判断目录和文件,如果是文件
			把文件的内容写入body
		*/
		//
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		p, _ := filepath.Rel(dir, name)
		header.Name = filepath.Join(basedir, p)
		//
		err = tr.WriteHeader(header)
		if err != nil {
			return err
		}
		//
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(name)
		if err != nil {
			return err
		}
		defer f.Close()
		//
		_, err = io.Copy(tr, f)
		if err != nil {
			return err
		}
		return nil
	})

	return nil
}

/*
func main() {
	l := "http://daily.zhihu.com/"
	//l := "http://pic.netbian.com/4kmingxing/index.html"
	//	http://59.110.12.72:7070/golang-spider/img.html
	//url := os.Args[1]
	urls, err := fetch(l)
	if err != nil {
		log.Fatal(err)
	}

	tmpdir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tmpdir)
	defer os.RemoveAll(tmpdir)
	err = downloadImags(urls, tmpdir)
	if err != nil {
		log.Panic(err)
	}
	//打包
	//tarfile(tmpdir)
	tarname := filepath.Join(".", "img.tar.gz")
	ff, err := os.Create(tarname)
	if err != nil {
		log.Fatal(err)
	}
	err = maketar(tmpdir, ff)
	if err != nil{
		log.Panic(err)
	}

}
*/

func fetchImages(w io.Writer, url string) {
	urls, err := fetch(url)
	if err != nil {
		log.Panic(err)
	}
	urls = cleanUrls(url, urls)
	tmpdir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Panic(err)
	}
	defer os.RemoveAll(tmpdir)
	err = downloadImags(urls, tmpdir)
	if err != nil {
		log.Panic(err)
	}
	//打包
	//tarfile(tmpdir)
	tarname := filepath.Join(".", "img.tar.gz")
	ff, err := os.Create(tarname)
	if err != nil {
		log.Fatal(err)
	}
	err = maketar(tmpdir, ff)
	if err != nil {
		log.Panic(err)
	}
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fetchImages(w, r.FormValue("u"))
}

func main() {
	http.HandleFunc("/", handleHTTP)
	http.ListenAndServe(":8000", nil)
}
