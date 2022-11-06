# World Development Indicators Data Prep

This project takes the default format of csv data from the World Development
Indicators database and converts the values in the `Series Name` column to
columns and the year columns to rows. This makes the data much easier to use
in programs like [Tableau](https://www.tableau.com/) without having to do a lot
of preparation work. See the following two tables for an example.

## Example Input

| Country Name   | Country Code | Series Name           | Series Code       | 2020 [YR2020]     | 2021 [YR2021]    |
| -------------- | ------------ | --------------------- | ----------------- | ----------------- | ---------------- |
| United States  | USA          | GDP growth (annual %) | NY.GDP.MKTP.KD.ZG | -3.40458965156256 | 5.67110719074167 |
| United States  | USA          | GDP (current US$)     | NY.GDP.MKTP.CD    | 20893743833000    | 22996100000000   |
| United Kingdom | GBR          | GDP growth (annual %) | NY.GDP.MKTP.KD.ZG | -9.27041083453351 | 7.44127309273938 |
| United Kingdom | GBR          | GDP (current US$)     | NY.GDP.MKTP.CD    | 2756900214107.32  | 3186859739185.02 |
| Canada         | CAN          | GDP growth (annual %) | NY.GDP.MKTP.KD.ZG | -5.23302430280606 | 4.56289386263273 |
| Canada         | CAN          | GDP (current US$)     | NY.GDP.MKTP.CD    | 1645423407568.36  | 1990761609665.23 |

## Example Output

| Country Name   | Country Code | Year | GDP growth (annual %) | GDP (current US$) |
| -------------- | ------------ | ---- | --------------------- | ----------------- |
| United States  | USA          | 2021 | 5.67110719074167      | 22996100000000    |
| United States  | USA          | 2020 | -3.40458965156256     | 20893743833000    |
| United Kingdom | GBR          | 2021 | 7.44127309273938      | 3186859739185.02  |
| United Kingdom | GBR          | 2020 | -9.27041083453351     | 2756900214107.32  |
| Canada         | CAN          | 2021 | 4.56289386263273      | 1990761609665.23  |
| Canada         | CAN          | 2020 | -5.23302430280606     | 1645423407568.36  |

## How to Use

1. [Install Go](https://go.dev/dl/) if you do not already have it installed.
2. Run `go run . [input path] [output path]`
