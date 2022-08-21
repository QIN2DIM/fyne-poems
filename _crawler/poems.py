import os
from collections import Counter
from urllib.request import getproxies

import requests
from bs4 import BeautifulSoup

PATH_DATASET = r"corpus.poems.txt"


def claim_dataset():
    """
    采集 500 條有關<愛情>的古詩文樣本
    :return:
    """
    corpus_memory = []
    url_entry = "https://so.gushiwen.cn/mingjus/default.aspx?tstr=%E7%88%B1%E6%83%85&astr=&cstr=&xstr="
    headers = {
        "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) "
                      "Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63"
    }
    for page in range(1, 11):
        url = f"{url_entry}&page={page}"
        print("_driver context - page={} url={}".format(page, url))

        resp = requests.get(url, headers=headers, proxies=getproxies())
        soup = BeautifulSoup(resp.text, "html.parser")
        cont = soup.find("div", class_="sons").find_all("div", class_="cont")
        for p in cont:
            corpus = p.text.strip().split("\n")
            if len(corpus) == 2:
                corpus_string = " ".join(corpus)
                corpus_memory.append(corpus_string)
            else:
                print(corpus, url)

    with open(PATH_DATASET, "w", encoding="utf8") as file:
        for corpus_string in corpus_memory:
            file.write(f"{corpus_string}\n")


def filter_dataset():
    if not os.path.exists(PATH_DATASET):
        return

    with open(PATH_DATASET, "r", encoding="utf8") as file:
        data = [i for i in file.read().split("\n") if i]

    # 統計語料長度 key：語料長度 value：該長度重複出現的次數
    print(Counter([len(i) for i in data]))

    s = sorted(data, key=lambda x: len(x))
    xc = [i for i in s if len(i) <= 30]

    # 選擇長度31以内的語料進行渲染
    print(len(xc))


def format_dataset():
    """
    將corpus格式化成 golang 字符串切片的格式，方便通信
    :return:
    """
    if not os.path.exists(PATH_DATASET):
        return

    with open(PATH_DATASET, "r", encoding="utf8") as file:
        data = [i for i in file.read().split("\n") if i]

    content = "{\n" + "".join([f"\t\"{i}\",\n" for i in data if len(i) <= 30]) + "}"
    box = f"var corpus = []string{content}\n"

    path_corpus = f"_{PATH_DATASET}"
    with open(path_corpus, "w", encoding="utf8") as file:
        file.write(box)


def demo():
    claim_dataset()
    filter_dataset()
    format_dataset()
