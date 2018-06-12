# Scala version must be 2.11

[Example: wikiedits.WikipediaAnalysis](https://ci.apache.org/projects/flink/flink-docs-release-1.5/quickstart/run_example_quickstart.html#writing-a-flink-program)

```pom
<dependencies>
    <dependency>
        <groupId>org.apache.flink</groupId>
        <artifactId>flink-java</artifactId>
        <version>${flink.version}</version>
    </dependency>
    <dependency>
        <groupId>org.apache.flink</groupId>
        <artifactId>flink-streaming-java_2.11</artifactId>
        <version>${flink.version}</version>
    </dependency>
    <dependency>
        <groupId>org.apache.flink</groupId>
        <artifactId>flink-clients_2.11</artifactId>
        <version>${flink.version}</version>
    </dependency>
    <dependency>
        <groupId>org.apache.flink</groupId>
        <artifactId>flink-connector-wikiedits_2.11</artifactId>
        <version>${flink.version}</version>
    </dependency>
</dependencies>
```

Notice: 2.11 must accord with Scala version. The latest version of Scala is 2.12. Must install 2.11

```bash
brew install scala@2.11
```