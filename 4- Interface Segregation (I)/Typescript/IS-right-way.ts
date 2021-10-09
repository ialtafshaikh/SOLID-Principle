// The following example shows an alternative approach that groups methods into more specific interfaces. 
// It describe a number of contracts that could be implemented individually by a simple printer or simple 
// copier or by a super printer:


interface Printer {
    printDocument(document: Document);
}


interface Stapler {
    stapleDocument(document: Document, tray: number);
}


interface Copier {
    copyDocument();
}

class SimplePrinter implements Printer {
    public printDocument(document: Document) {
        //...
    }
}