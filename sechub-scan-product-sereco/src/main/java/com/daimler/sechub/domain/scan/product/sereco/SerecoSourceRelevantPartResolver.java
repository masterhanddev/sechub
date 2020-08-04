package com.daimler.sechub.domain.scan.product.sereco;

import java.util.regex.Pattern;

import org.springframework.stereotype.Component;

@Component
public class SerecoSourceRelevantPartResolver{
    
    
    private static final Pattern P = Pattern.compile("\\s");

    public String toRelevantPart(String source) {
        String result = P.matcher(source).replaceAll("");
        return result.toLowerCase();
    }
}
