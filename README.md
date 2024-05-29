Exemple de programme Golang pour récupérer les caractéristiques d'un livre
Go
// Projet  sur https://console.cloud.google.com/projectcreate?previousPage=%2Fprojectselector2%2Fapis%2Fdashboard%3Fhl%3Dfr%26supportedpurview%3Dproject&hl=fr&supportedpurview=project
Nom du projet: 'Htitus2024'
lien: http://onvaessayer.org/appinventor/baseApps/Scan_isbn.php
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

// Structure pour stocker les informations du livre
type Book struct {
    ID           string `json:"id"`
    VolumeInfo struct {
        Title       string `json:"title"`
        Authors     []string `json:"authors"`
        Description string `json:"description"`
        ImageLinks struct {
            Thumbnail string `json:"thumbnail"`
        } `json:"imageLinks"`
    } `json:"volumeInfo"`
}

func main() {
    // Définir l'URL de l'API
    url := "https://www.googleapis.com/books/v1/volumes?q=isbn:9780245505836"

    // Créer une requête HTTP
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Définir le contexte de la requête
    ctx := context.Background()

    // Envoyer la requête et obtenir la réponse
    resp, err := http.DefaultClient.Do(req.WithContext(ctx))
    if err != nil {
        fmt.Println(err)
        return
    }

    // Décoder le corps de la réponse en JSON
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Fermer le corps de la réponse
    resp.Body.Close()

    // Définir une variable pour stocker le livre
    var book Book

    // Décoder le JSON dans la structure
    err = json.Unmarshal(body, &book)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Afficher les informations du livre
    fmt.Println("Titre:", book.VolumeInfo.Title)
    fmt.Println("Auteurs:", strings.Join(book.VolumeInfo.Authors, ", "))
    fmt.Println("Description:", book.VolumeInfo.Description)
    fmt.Println("Lien vers l'image miniature:", book.VolumeInfo.ImageLinks.Thumbnail)
}

Utilisez ce code avec précaution.

Ce programme effectue les actions suivantes :

    Définir l'URL de l'API: L'URL de l'API est définie avec l'ISBN du livre que vous souhaitez rechercher.
    Créer une requête HTTP: Une requête HTTP GET est créée avec l'URL de l'API.
    Définir le contexte de la requête: Le contexte de la requête est défini.
    Envoyer la requête et obtenir la réponse: La requête est envoyée et la réponse est reçue.
    Décoder le corps de la réponse en JSON: Le corps de la réponse est décodé en JSON.
    Définir une variable pour stocker le livre: Une variable est définie pour stocker les informations du livre.
    Décoder le JSON dans la structure: Le JSON est décodé dans la structure Book.
    Afficher les informations du livre: Les informations du livre sont affichées sur la console.

Cet exemple montre comment utiliser l'API Google Books pour récupérer les caractéristiques d'un livre.

Voici quelques points à noter :

    Vous devez remplacer l'ISBN dans l'URL par l'ISBN du livre que vous souhaitez rechercher.
    Vous pouvez modifier le programme pour afficher d'autres informations du livre.
    Vous pouvez également utiliser l'API Google Books pour rechercher des livres par d'autres critères, tels que le titre, l'auteur ou le sujet.

Pour plus d'informations, veuillez consulter la documentation de l'API Google Books :

