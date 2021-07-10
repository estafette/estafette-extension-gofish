# extensions/gofish

This extension helps with creating [GoFish](https://gofi.sh/) food to store in a rig repository

## Parameters

| Parameter   | Type   | Values                                                                  |
| ----------- | ------ | ----------------------------------------------------------------------- |
| description | string | The description of what the application in the food does             |

## Usage

In order to use this extension in your `.estafette.yaml` manifest use the following snippets:

### Linting

```yaml
  create-gofish-food:
    image: extensions/gofish:stable
    description: This is my awesome own food that does abc
```
