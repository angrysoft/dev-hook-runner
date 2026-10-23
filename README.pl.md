# dhr — Dev Hook Runner

[English](README.md)

`dhr` uruchamia hooki zdefiniowane w plikach YAML. Hooki mogą wykonywać polecenia i skrypty powłoki, korzystać ze zmiennych środowiskowych oraz deklarować zależności od innych hooków.

## Wymagania

- Go w wersji zgodnej z `go.mod` (1.27.1 lub nowsza)
- Bash, jeśli hook zawiera skrypt `script`

## Budowanie i testy

```sh
make build
make test
```

Plik wykonywalny zostanie utworzony jako `bin/dhr`. Możesz też uruchomić program bezpośrednio:

```sh
go run . [opcje] [nazwa-hooka]
```

Pozostałe dostępne cele Makefile:

```sh
make help
make install
make cross-build
make clean
```

`make install` instaluje program w `$GOPATH/bin`, a `make cross-build` przygotowuje pliki dla Linuxa, macOS i Windows.

## Konfiguracja hooków

Domyślnie `dhr` szuka hooków w katalogu `.hooks.d` bieżącego katalogu. Każdy hook powinien znajdować się w pliku o nazwie `<nazwa>.hook.yml`; do uruchomienia hooka podaje się tę nazwę bez rozszerzenia.

Przykładowa struktura:

```text
projekt/
├── .hooks.d/
│   ├── build.hook.yml
│   └── prepare.hook.yml
└── ...
```

Przykładowy plik `.hooks.d/build.hook.yml`:

```yaml
name: build
description: Buduje projekt
depends_on:
  - prepare
env_file: .hooks.d/build.env.yml
execs:
  - cmd: go
    args:
      - build
      - ./...
script: echo "Budowanie zakończone"
needs_args: false
```

Wszystkie pola są opcjonalne:

| Pole | Opis |
| --- | --- |
| `name` | Nazwa wyświetlana podczas uruchamiania hooka. |
| `description` | Krótki opis, wyświetlany przy listowaniu hooków. |
| `depends_on` | Lista nazw hooków do uruchomienia przed tym hookiem. |
| `env_file` | Ścieżka do pliku YAML ze zmiennymi środowiskowymi. |
| `execs` | Lista poleceń do uruchomienia. Każdy element zawiera `cmd` oraz listę argumentów `args`. |
| `script` | Skrypt uruchamiany przez `bash -c`. |
| `needs_args` | Gdy ma wartość `true`, hook wymaga argumentu przekazanego opcją `--arg` lub `-a`; wartość jest dostępna jako `HOOK_ARG`. |

Polecenia z `execs` są wykonywane po kolei, a następnie uruchamiany jest `script`. Polecenia i skrypt dziedziczą środowisko procesu oraz otrzymują zmienne z `env_file`.

Przykładowy plik środowiskowy `.hooks.d/build.env.yml`:

```yaml
vars:
  - key: BUILD_MODE
    value: release
```

## Użycie

Uruchomienie hooka z bieżącego katalogu:

```sh
dhr build
```

Przekazanie argumentu do hooka wymagającego argumentów:

```sh
dhr --arg release build
```

Wskazanie katalogu projektu (najlepiej ścieżką bezwzględną):

```sh
dhr --dir /ścieżka/do/projektu build
```

Wyświetlenie hooków dostępnych w katalogu:

```sh
dhr --list
dhr --dir /ścieżka/do/projektu --list
```

| Opcja | Skrót | Opis |
| --- | --- | --- |
| `--dir` | `-d` | Katalog projektu zawierający `.hooks.d`; domyślnie `.`. |
| `--list` | `-l` | Wyświetla nazwy i opisy hooków. |
| `--arg` | `-a` | Argument przekazywany do hooka. |

Zależności są uruchamiane przed hookiem, który je deklaruje. Nazwy w `depends_on` odpowiadają nazwom plików bez przyrostka `.hook.yml`.
