f = open("./input.txt", "r")
batch = f.read()
split = batch.split("\n\n")

elf_list = []

elf_enumerable = enumerate(split)
for (i, items) in elf_enumerable:
    arr = [int(x) for x in items.removesuffix("\n").split("\n")]
    elf_list.append(
        {"elf_number": i, "meals": arr, "calories": sum(arr)})


def sortElfFunction(elf):
    return elf['calories']


elf_list.sort(key=sortElfFunction, reverse=True)

print("\n\n =======================================================================")
print("|----------------------------ELF-LEADERBOARD----------------------------|")
print(" =======================================================================")
for x in elf_list[0:3]:
    print(x)
print(" =======================================================================\n\n")

total = 0
for x in elf_list[0:3]:
    total += x['calories']

print("Total calorie count: {}".format(total))
