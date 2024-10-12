import requests

def main(website):
    # 打开txt文件
    with open('address.txt', 'r') as f:
        lines = f.readlines()

    # 逐行读取每行内容，拼接URL并发送GET请求
    for line in lines:
        # 去除每行末尾的换行符
        line = line.strip()
        # 拼接URL，将每行内容作为参数添加到URL中
        url = website+ '/api/nft/v1/account/own/' + line + '/contract'
        # 发送GET请求
        response = requests.get(url)

if __name__ == '__main__':
    website = 'http://decert.me/'
    main(website)