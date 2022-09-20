package category

import (
    "gohub/pkg/database"
)

func Get(idstr string) (category Category) {
    database.DB.Where("id", idstr).First(&category)
    return
}

func GetBy(field, value string) (category Category) {
    database.DB.Where("? = ?", field, value).First(&category)
    return
}

func All() (categories []Category) {
    database.DB.Find(&categories)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Category{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}