
with open("./color.go","r",encoding="utf-8") as f:
  lines = f.readlines()

  for index,line in enumerate(lines):
    s = line.strip()
    if s.find("//") == 0:
        s = s[3:] # 描述
        next = lines[index + 1]
        next = next[5:next.index("(")]
        print(f"logger.{next}()")


