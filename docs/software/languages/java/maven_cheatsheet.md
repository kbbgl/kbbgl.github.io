# Maven Cheatsheet

## Clean

Clean the maven project by deleting the target directory:

```bash
mvn clean
```

## Install Dependencies

This command builds the maven project and installs the project files (JAR, WAR, pom.xml, etc) to the local repository. Skips tests:

```bash
mvn install -DskipTests [-f "/Users/kobbi.gal/development/be-services/be/services/build/pom.xml"]
```

## Build

This command builds the maven project and packages them into a JAR, WAR, etc

```bash
mvn package -DskipTests
```

## Compile

```bash
mvn compile
```
